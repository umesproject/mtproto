package mtproto

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"

	"github.com/umesproject/mtproto/internal/encoding/tl"
	"github.com/umesproject/mtproto/internal/mode"
	"github.com/umesproject/mtproto/internal/mtproto/messages"
	"github.com/umesproject/mtproto/internal/mtproto/objects"
	"github.com/umesproject/mtproto/internal/session"
	"github.com/umesproject/mtproto/internal/transport"
	"github.com/umesproject/mtproto/internal/utils"
)

type MTProto struct {
	debug        bool
	addr         string
	transport    transport.Transport
	stopRoutines context.CancelFunc // stopping ping, read, etc. routines
	routineswg   sync.WaitGroup     // WaitGroup for being sure that all routines are stopped

	// ключ авторизации. изменять можно только через setAuthKey
	authKey []byte

	// хеш ключа авторизации. изменять можно только через setAuthKey
	authKeyHash []byte

	// соль сессии
	serverSalt int64
	encrypted  bool
	sessionId  int64

	// общий мьютекс
	mutex sync.Mutex

	// каналы, которые ожидают ответа rpc. ответ записывается в канал и удаляется
	responseChannels *utils.SyncIntObjectChan
	expectedTypes    *utils.SyncIntReflectTypes // uses for parcing bool values in rpc result for example

	// идентификаторы сообщений, нужны что бы посылать и принимать сообщения.
	seqNoMutex sync.Mutex
	seqNo      int32
	msgId      int64

	// айдишники DC для КОНКРЕТНОГО Приложения и клиента. Может меняться, но фиксирована для
	// связки приложение+клиент
	dclist map[int]string

	session *session.Session

	// один из публичных ключей telegram. нужен только для создания сессии.
	publicKey *rsa.PublicKey

	// serviceChannel нужен только на время создания ключей, т.к. это
	// не RpcResult, поэтому все данные отдаются в один поток без
	// привязки к MsgID
	serviceChannel       chan tl.Object
	serviceModeActivated bool

	//! DEPRECATED RecoverFunc используется только до того момента, когда из пакета будут убраны все паники
	RecoverFunc func(i any)
	// if set, all critical errors writing to this channel
	Warnings chan error

	serverRequestHandlers []customHandlerFunc
}

type customHandlerFunc = func(i any) bool

type Config struct {
	Session    *session.Session
	Debug      bool
	ServerHost string
	PublicKey  *rsa.PublicKey
	ProxyUrl   string
}

func NewMTProto(c Config) (*MTProto, error) {
	m := &MTProto{
		addr:                  c.ServerHost,
		encrypted:             c.Session != nil && len(c.Session.Key) > 0,
		sessionId:             utils.GenerateSessionID(),
		serviceChannel:        make(chan tl.Object),
		publicKey:             c.PublicKey,
		responseChannels:      utils.NewSyncIntObjectChan(),
		expectedTypes:         utils.NewSyncIntReflectTypes(),
		serverRequestHandlers: make([]customHandlerFunc, 0),
		dclist:                defaultDCList(),
		session:               c.Session,
	}

	if c.Session != nil && len(c.Session.Key) > 0 {
		m.LoadSession(c.Session)
	}

	return m, nil
}

func (m *MTProto) SetDCList(in map[int]string) {
	if m.dclist == nil {
		m.dclist = make(map[int]string)
	}
	for k, v := range in {
		m.dclist[k] = v
	}
}

func (m *MTProto) GetSessionJSON() string {
	// m.authKey = s.Key
	// 	m.authKeyHash = s.Hash
	// 	m.serverSalt = s.Salt
	// 	m.addr = s.Hostname

	s := session.Session{
		Key:      m.authKey,
		Hash:     m.authKeyHash,
		Salt:     m.serverSalt,
		Hostname: m.addr,
	}

	res, _ := json.Marshal(s)
	return string(res)
}

func (m *MTProto) CreateConnection() error {
	ctx, cancelfunc := context.WithCancel(context.Background())
	m.stopRoutines = cancelfunc

	err := m.connect(ctx)
	if err != nil {
		return err
	}

	// start reading responses from the server
	m.startReadingResponses(ctx)

	// get new authKey if need
	if !m.encrypted {
		err = m.makeAuthKey()
		if err != nil {
			return errors.Wrap(err, "making auth key")
		}
	}

	// start keepalive pinging
	m.startPinging(ctx)

	return nil
}

const defaultTimeout = 65 * time.Second // 60 seconds is maximum timeouts without pings

func (m *MTProto) connect(ctx context.Context) error {
	var err error
	m.transport, err = transport.NewTransport(
		m,
		transport.TCPConnConfig{
			Ctx:     ctx,
			Host:    m.addr,
			Timeout: defaultTimeout,
		},
		mode.Intermediate,
	)
	if err != nil {
		return errors.Wrap(err, "can't connect")
	}

	CloseOnCancel(ctx, m.transport)
	return nil
}

func (m *MTProto) makeRequest(data tl.Object, expectedTypes ...reflect.Type) (any, error) {
	resp, err := m.sendPacket(data, expectedTypes...)
	if err != nil {
		return nil, errors.Wrap(err, "sending message")
	}

	response := <-resp

	switch r := response.(type) {
	case *objects.RpcError:
		realErr := RpcErrorToNative(r)

		err = m.tryToProcessErr(realErr.(*ErrResponseCode))
		if err != nil {
			return nil, err
		}

		return m.makeRequest(data, expectedTypes...)

	case *errorSessionConfigsChanged:
		return m.makeRequest(data, expectedTypes...)

	}

	return tl.UnwrapNativeTypes(response), nil
}

// Disconnect is closing current TCP connection and stopping all routines like pinging, reading etc.
func (m *MTProto) Disconnect() error {
	// stop all routines
	m.stopRoutines()

	// TODO: close ALL CHANNELS

	return nil
}

func (m *MTProto) Reconnect(makeAuthKeyAgain bool) error {
	if makeAuthKeyAgain {
		m.encrypted = false
	}

	err := m.Disconnect()
	if err != nil {
		return errors.Wrap(err, "disconnecting")
	}

	err = m.CreateConnection()
	return errors.Wrap(err, "recreating connection")
}

// StartPinging pings the server that everything is fine, the client is online
// you just need to run and forget about it
func (m *MTProto) startPinging(ctx context.Context) {
	m.routineswg.Add(1)

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		defer m.routineswg.Done()
		defer func() {
			fmt.Println("StartPinging is exit!")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				//_, err := m.ping(0xCADACADA) //nolint:gomnd not magic
				_, err := m.ping_delay_disconnect(time.Now().UnixNano(), 75) //保持链接
				if err != nil {
					fmt.Println("ping unsuccsesful")
					//go func() {
					//	err = m.Reconnect()
					//	if err != nil {
					//		m.warnError(errors.Wrap(err, "can't reconnect"))
					//	}
					//}()
					//return
				} else {
					fmt.Println("ping succsesful")
				}
			}
		}
	}()
}
func (m *MTProto) startReadingResponses(ctx context.Context) {
	m.routineswg.Add(1)
	go func() {
		defer m.routineswg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := m.readMsg()
				switch err {
				case nil: // skip
				case context.Canceled:
					return
				case io.EOF:
					err = m.Reconnect(false)
					if err != nil {
						m.warnError(errors.Wrap(err, "can't reconnect"))
					}
				default:
					check(err)
				}
			}
		}
	}()
}

func (m *MTProto) readMsg() error {
	if m.transport == nil {
		return errors.New("must setup connection before reading messages")
	}

	response, err := m.transport.ReadMsg()
	if err != nil {
		if e, ok := err.(transport.ErrCode); ok {
			return &ErrResponseCode{Code: int(e)}
		}
		switch err {
		case io.EOF, context.Canceled:
			return err
		default:
			return errors.Wrap(err, "reading message")
		}
	}

	if m.serviceModeActivated {
		var obj tl.Object
		// сервисные сообщения ГАРАНТИРОВАННО в теле содержат TL.
		obj, err = tl.DecodeUnknownObject(response.GetMsg())
		if err != nil {
			return errors.Wrap(err, "parsing object")
		}
		m.serviceChannel <- obj
		return nil
	}

	err = m.processResponse(response)
	if err != nil {
		return errors.Wrap(err, "processing response")
	}
	return nil
}

func (m *MTProto) processResponse(msg messages.Common) error {
	var data tl.Object
	var err error
	if et, ok := m.expectedTypes.Get(msg.GetMsgID()); ok && len(et) > 0 {
		data, err = tl.DecodeUnknownObject(msg.GetMsg(), et...)
	} else {
		data, err = tl.DecodeUnknownObject(msg.GetMsg())
	}
	if err != nil {
		return errors.Wrap(err, "unmarshaling response")
	}

messageTypeSwitching:
	switch message := data.(type) {
	case *objects.MessageContainer:
		for _, v := range *message {
			err := m.processResponse(v)
			if err != nil {
				return errors.Wrap(err, "processing item in container")
			}
		}

	case *objects.BadServerSalt:
		m.serverSalt = message.NewSalt

		m.mutex.Lock()
		for _, k := range m.responseChannels.Keys() {
			v, _ := m.responseChannels.Get(k)
			v <- &errorSessionConfigsChanged{}
		}
		m.mutex.Unlock()

	case *objects.NewSessionCreated:
		m.serverSalt = message.ServerSalt
	case *objects.Pong, *objects.MsgsAck:
		// игнорим, пришло и пришло, че бубнить то

	case *objects.BadMsgNotification:
		pp.Println(message)
		panic(message) // for debug, looks like this message is important
		return BadMsgErrorFromNative(message)

	case *objects.RpcResult:
		obj := message.Obj
		if v, ok := obj.(*objects.GzipPacked); ok {
			obj = v.Obj
		}

		err := m.writeRPCResponse(int(message.ReqMsgID), obj)
		if err != nil {
			return errors.Wrap(err, "writing RPC response")
		}

	case *objects.GzipPacked:
		// sometimes telegram server returns gzip for unknown reason. so, we are extracting data from gzip and
		// reprocess it again
		data = message.Obj
		goto messageTypeSwitching

	default:
		processed := false
		for _, f := range m.serverRequestHandlers {
			processed = f(message)
			if processed {
				break
			}
		}
		if !processed {
			m.warnError(errors.New("got nonsystem message from server: " + reflect.TypeOf(message).String()))
		}
	}

	if (msg.GetSeqNo() & 1) != 0 {
		_, err := m.MakeRequest(&objects.MsgsAck{MsgIDs: []int64{int64(msg.GetMsgID())}})
		if err != nil {
			return errors.Wrap(err, "sending ack")
		}
	}

	return nil
}

// tryToProcessErr пытается автоматически решить ошибку полученную от сервера. в случае успеха вернет nil,
// в случае если нет способа решить эту проблему, возвращается сама ошибка
// если в процессе решения появлиась еще одна ошибка, то она оборачивается в errors.Wrap, основная
// игнорируется (потому что гарантируется, что обработка ошибки надежна, и параллельная ошибка это что-то из
// ряда вон выходящее)
func (m *MTProto) tryToProcessErr(e *ErrResponseCode) error {
	switch e.Message {
	case "PHONE_MIGRATE_X":
		return m.ConnectAgainToDC(e.AdditionalInfo.(int))
	case "FILE_MIGRATE_X":
		return e
	default:
		return e
	}
}

func (m *MTProto) DebugPrintf(format string, a ...interface{}) (n int, err error) {
	if m.debug {
		return fmt.Printf(format, a...)
	}

	return 0, nil
}

// Author: Kliton
// Reconnect to specified DC
func (m *MTProto) ConnectAgainToDC(dc int) error {
	newIP, found := m.dclist[dc]
	if !found {
		return errors.New(fmt.Sprint("dc", dc, "ip not found"))

	}

	log.Println("Connecting to [DC]", dc, "with [IP]", newIP)
	m.addr = newIP
	err := m.Reconnect(true)
	return err
}
