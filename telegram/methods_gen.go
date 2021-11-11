package telegram

import (
	"reflect"
	//"strings"
	"log"

	errors "github.com/pkg/errors"
)

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountAcceptAuthorizationParams struct {
	BotID       int32                       // Bot ID
	Scope       string                      // Telegram Passport element types requested by the service
	PublicKey   string                      // Service's public key
	ValueHashes []*SecureValueHash          // Types of values sent and their hashes
	Credentials *SecureCredentialsEncrypted // Encrypted values
}

func (*AccountAcceptAuthorizationParams) CRC() uint32 {
	return 0xe7027c94
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountAcceptAuthorization(botID int32, scope, publicKey string, valueHashes []*SecureValueHash, credentials *SecureCredentialsEncrypted) (bool, error) {
	responseData, err := c.MakeRequest(&AccountAcceptAuthorizationParams{
		BotID:       botID,
		Credentials: credentials,
		PublicKey:   publicKey,
		Scope:       scope,
		ValueHashes: valueHashes,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountAcceptAuthorization")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountCancelPasswordEmailParams struct{}

func (*AccountCancelPasswordEmailParams) CRC() uint32 {
	return 0xc1cbd5b6
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountCancelPasswordEmail() (bool, error) {
	responseData, err := c.MakeRequest(&AccountCancelPasswordEmailParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountCancelPasswordEmail")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountChangePhoneParams struct {
	PhoneNumber   string // Phone number in international format
	PhoneCodeHash string // SMS-message ID
	PhoneCode     string
}

func (*AccountChangePhoneParams) CRC() uint32 {
	return 0x70c32edb
}

// Registers a validated phone number in the system.
func (c *Client) AccountChangePhone(phoneNumber, phoneCodeHash, phoneCode string) (User, error) {
	responseData, err := c.MakeRequest(&AccountChangePhoneParams{
		PhoneCode:     phoneCode,
		PhoneCodeHash: phoneCodeHash,
		PhoneNumber:   phoneNumber,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountChangePhone")
	}

	resp, ok := responseData.(User)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountCheckUsernameParams struct {
	Username string
}

func (*AccountCheckUsernameParams) CRC() uint32 {
	return 0x2714d86c
}

// Registers a validated phone number in the system.
func (c *Client) AccountCheckUsername(username string) (bool, error) {
	responseData, err := c.MakeRequest(&AccountCheckUsernameParams{Username: username})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountCheckUsername")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountConfirmPasswordEmailParams struct {
	Code string
}

func (*AccountConfirmPasswordEmailParams) CRC() uint32 {
	return 0x8fdf1920
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountConfirmPasswordEmail(code string) (bool, error) {
	responseData, err := c.MakeRequest(&AccountConfirmPasswordEmailParams{Code: code})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountConfirmPasswordEmail")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountConfirmPhoneParams struct {
	PhoneCodeHash string // SMS-message ID
	PhoneCode     string
}

func (*AccountConfirmPhoneParams) CRC() uint32 {
	return 0x5f2178c3
}

// Registers a validated phone number in the system.
func (c *Client) AccountConfirmPhone(phoneCodeHash, phoneCode string) (bool, error) {
	responseData, err := c.MakeRequest(&AccountConfirmPhoneParams{
		PhoneCode:     phoneCode,
		PhoneCodeHash: phoneCodeHash,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountConfirmPhone")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountCreateThemeParams struct {
	Slug     string
	Title    string
	Document InputDocument       `tl:"flag:2"`
	Settings *InputThemeSettings `tl:"flag:3"` // Settings for the code type to send
}

func (*AccountCreateThemeParams) CRC() uint32 {
	return 0x8432c21f
}

func (*AccountCreateThemeParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountCreateTheme(slug, title string, document InputDocument, settings *InputThemeSettings) (*Theme, error) {
	responseData, err := c.MakeRequest(&AccountCreateThemeParams{
		Document: document,
		Settings: settings,
		Slug:     slug,
		Title:    title,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountCreateTheme")
	}

	resp, ok := responseData.(*Theme)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountDeleteAccountParams struct {
	Reason string
}

func (*AccountDeleteAccountParams) CRC() uint32 {
	return 0x418d4e0b
}

// Registers a validated phone number in the system.
func (c *Client) AccountDeleteAccount(reason string) (bool, error) {
	responseData, err := c.MakeRequest(&AccountDeleteAccountParams{Reason: reason})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountDeleteAccount")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountDeleteSecureValueParams struct {
	Types []SecureValueType
}

func (*AccountDeleteSecureValueParams) CRC() uint32 {
	return 0xb880bc4b
}

// Registers a validated phone number in the system.
func (c *Client) AccountDeleteSecureValue(types []SecureValueType) (bool, error) {
	responseData, err := c.MakeRequest(&AccountDeleteSecureValueParams{Types: types})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountDeleteSecureValue")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountFinishTakeoutSessionParams struct {
	Success bool `tl:"flag:0,encoded_in_bitflags"`
}

func (*AccountFinishTakeoutSessionParams) CRC() uint32 {
	return 0x1d2652ee
}

func (*AccountFinishTakeoutSessionParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountFinishTakeoutSession(success bool) (bool, error) {
	responseData, err := c.MakeRequest(&AccountFinishTakeoutSessionParams{Success: success})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountFinishTakeoutSession")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetAccountTtlParams struct{}

func (*AccountGetAccountTtlParams) CRC() uint32 {
	return 0x8fc711d
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetAccountTtl() (*AccountDaysTtl, error) {
	responseData, err := c.MakeRequest(&AccountGetAccountTtlParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetAccountTtl")
	}

	resp, ok := responseData.(*AccountDaysTtl)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetAllSecureValuesParams struct{}

func (*AccountGetAllSecureValuesParams) CRC() uint32 {
	return 0xb288bc7d
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetAllSecureValues() ([]*SecureValue, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&AccountGetAllSecureValuesParams{}, reflect.TypeOf([]*SecureValue{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetAllSecureValues")
	}

	resp, ok := responseData.([]*SecureValue)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetAuthorizationFormParams struct {
	BotID     int32
	Scope     string
	PublicKey string
}

func (*AccountGetAuthorizationFormParams) CRC() uint32 {
	return 0xb86ba8e1
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetAuthorizationForm(botID int32, scope, publicKey string) (*AccountAuthorizationForm, error) {
	responseData, err := c.MakeRequest(&AccountGetAuthorizationFormParams{
		BotID:     botID,
		PublicKey: publicKey,
		Scope:     scope,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetAuthorizationForm")
	}

	resp, ok := responseData.(*AccountAuthorizationForm)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetAuthorizationsParams struct{}

func (*AccountGetAuthorizationsParams) CRC() uint32 {
	return 0xe320c158
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetAuthorizations() (*AccountAuthorizations, error) {
	responseData, err := c.MakeRequest(&AccountGetAuthorizationsParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetAuthorizations")
	}

	resp, ok := responseData.(*AccountAuthorizations)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetAutoDownloadSettingsParams struct{}

func (*AccountGetAutoDownloadSettingsParams) CRC() uint32 {
	return 0x56da0b3f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetAutoDownloadSettings() (*AccountAutoDownloadSettings, error) {
	responseData, err := c.MakeRequest(&AccountGetAutoDownloadSettingsParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetAutoDownloadSettings")
	}

	resp, ok := responseData.(*AccountAutoDownloadSettings)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetContactSignUpNotificationParams struct{}

func (*AccountGetContactSignUpNotificationParams) CRC() uint32 {
	return 0x9f07c728
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetContactSignUpNotification() (bool, error) {
	responseData, err := c.MakeRequest(&AccountGetContactSignUpNotificationParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountGetContactSignUpNotification")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetContentSettingsParams struct{}

func (*AccountGetContentSettingsParams) CRC() uint32 {
	return 0x8b9b4dae
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetContentSettings() (*AccountContentSettings, error) {
	responseData, err := c.MakeRequest(&AccountGetContentSettingsParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetContentSettings")
	}

	resp, ok := responseData.(*AccountContentSettings)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetGlobalPrivacySettingsParams struct{}

func (*AccountGetGlobalPrivacySettingsParams) CRC() uint32 {
	return 0xeb2b4cf6
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetGlobalPrivacySettings() (*GlobalPrivacySettings, error) {
	responseData, err := c.MakeRequest(&AccountGetGlobalPrivacySettingsParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetGlobalPrivacySettings")
	}

	resp, ok := responseData.(*GlobalPrivacySettings)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetMultiWallPapersParams struct {
	Wallpapers []InputWallPaper
}

func (*AccountGetMultiWallPapersParams) CRC() uint32 {
	return 0x65ad71dc
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetMultiWallPapers(wallpapers []InputWallPaper) ([]WallPaper, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&AccountGetMultiWallPapersParams{Wallpapers: wallpapers}, reflect.TypeOf([]WallPaper{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetMultiWallPapers")
	}

	resp, ok := responseData.([]WallPaper)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetNotifyExceptionsParams struct {
	CompareSound bool            `tl:"flag:1,encoded_in_bitflags"`
	Peer         InputNotifyPeer `tl:"flag:0"`
}

func (*AccountGetNotifyExceptionsParams) CRC() uint32 {
	return 0x53577479
}

func (*AccountGetNotifyExceptionsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetNotifyExceptions(compareSound bool, peer InputNotifyPeer) (Updates, error) {
	responseData, err := c.MakeRequest(&AccountGetNotifyExceptionsParams{
		CompareSound: compareSound,
		Peer:         peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetNotifyExceptions")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetNotifySettingsParams struct {
	Peer InputNotifyPeer
}

func (*AccountGetNotifySettingsParams) CRC() uint32 {
	return 0x12b3ad31
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetNotifySettings(peer InputNotifyPeer) (*PeerNotifySettings, error) {
	responseData, err := c.MakeRequest(&AccountGetNotifySettingsParams{Peer: peer})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetNotifySettings")
	}

	resp, ok := responseData.(*PeerNotifySettings)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetPasswordParams struct{}

func (*AccountGetPasswordParams) CRC() uint32 {
	return 0x548a30f5
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetPassword() (*AccountPassword, error) {
	responseData, err := c.MakeRequest(&AccountGetPasswordParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetPassword")
	}

	resp, ok := responseData.(*AccountPassword)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetPasswordSettingsParams struct {
	Password InputCheckPasswordSRP
}

func (*AccountGetPasswordSettingsParams) CRC() uint32 {
	return 0x9cd4eaf9
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetPasswordSettings(password InputCheckPasswordSRP) (*AccountPasswordSettings, error) {
	responseData, err := c.MakeRequest(&AccountGetPasswordSettingsParams{Password: password})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetPasswordSettings")
	}

	resp, ok := responseData.(*AccountPasswordSettings)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetPrivacyParams struct {
	Key InputPrivacyKey
}

func (*AccountGetPrivacyParams) CRC() uint32 {
	return 0xdadbc950
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetPrivacy(key InputPrivacyKey) (*AccountPrivacyRules, error) {
	responseData, err := c.MakeRequest(&AccountGetPrivacyParams{Key: key})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetPrivacy")
	}

	resp, ok := responseData.(*AccountPrivacyRules)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetSecureValueParams struct {
	Types []SecureValueType
}

func (*AccountGetSecureValueParams) CRC() uint32 {
	return 0x73665bc2
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetSecureValue(types []SecureValueType) ([]*SecureValue, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&AccountGetSecureValueParams{Types: types}, reflect.TypeOf([]*SecureValue{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetSecureValue")
	}

	resp, ok := responseData.([]*SecureValue)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetThemeParams struct {
	Format     string
	Theme      InputTheme
	DocumentID int64
}

func (*AccountGetThemeParams) CRC() uint32 {
	return 0x8d9d742b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetTheme(format string, theme InputTheme, documentID int64) (*Theme, error) {
	responseData, err := c.MakeRequest(&AccountGetThemeParams{
		DocumentID: documentID,
		Format:     format,
		Theme:      theme,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetTheme")
	}

	resp, ok := responseData.(*Theme)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetThemesParams struct {
	Format string
	Hash   int32
}

func (*AccountGetThemesParams) CRC() uint32 {
	return 0x285946f8
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetThemes(format string, hash int32) (AccountThemes, error) {
	responseData, err := c.MakeRequest(&AccountGetThemesParams{
		Format: format,
		Hash:   hash,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetThemes")
	}

	resp, ok := responseData.(AccountThemes)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetTmpPasswordParams struct {
	Password InputCheckPasswordSRP
	Period   int32
}

func (*AccountGetTmpPasswordParams) CRC() uint32 {
	return 0x449e0b51
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetTmpPassword(password InputCheckPasswordSRP, period int32) (*AccountTmpPassword, error) {
	responseData, err := c.MakeRequest(&AccountGetTmpPasswordParams{
		Password: password,
		Period:   period,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetTmpPassword")
	}

	resp, ok := responseData.(*AccountTmpPassword)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountGetWallPaperParams struct {
	Wallpaper InputWallPaper
}

func (*AccountGetWallPaperParams) CRC() uint32 {
	return 0xfc8ddbea
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountGetWallPaper(wallpaper InputWallPaper) (WallPaper, error) {
	responseData, err := c.MakeRequest(&AccountGetWallPaperParams{Wallpaper: wallpaper})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetWallPaper")
	}

	resp, ok := responseData.(WallPaper)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetWallPapersParams struct {
	Hash int32
}

func (*AccountGetWallPapersParams) CRC() uint32 {
	return 0xaabb1763
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetWallPapers(hash int32) (AccountWallPapers, error) {
	responseData, err := c.MakeRequest(&AccountGetWallPapersParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetWallPapers")
	}

	resp, ok := responseData.(AccountWallPapers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountGetWebAuthorizationsParams struct{}

func (*AccountGetWebAuthorizationsParams) CRC() uint32 {
	return 0x182e6d6f
}

// Registers a validated phone number in the system.
func (c *Client) AccountGetWebAuthorizations() (*AccountWebAuthorizations, error) {
	responseData, err := c.MakeRequest(&AccountGetWebAuthorizationsParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountGetWebAuthorizations")
	}

	resp, ok := responseData.(*AccountWebAuthorizations)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountInitTakeoutSessionParams struct {
	Contacts          bool  `tl:"flag:0,encoded_in_bitflags"`
	MessageUsers      bool  `tl:"flag:1,encoded_in_bitflags"`
	MessageChats      bool  `tl:"flag:2,encoded_in_bitflags"`
	MessageMegagroups bool  `tl:"flag:3,encoded_in_bitflags"`
	MessageChannels   bool  `tl:"flag:4,encoded_in_bitflags"`
	Files             bool  `tl:"flag:5,encoded_in_bitflags"`
	FileMaxSize       int32 `tl:"flag:5"`
}

func (*AccountInitTakeoutSessionParams) CRC() uint32 {
	return 0xf05b4804
}

func (*AccountInitTakeoutSessionParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountInitTakeoutSession(params *AccountInitTakeoutSessionParams) (*AccountTakeout, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountInitTakeoutSession")
	}

	resp, ok := responseData.(*AccountTakeout)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountInstallThemeParams struct {
	Dark   bool       `tl:"flag:0,encoded_in_bitflags"`
	Format string     `tl:"flag:1"`
	Theme  InputTheme `tl:"flag:1"`
}

func (*AccountInstallThemeParams) CRC() uint32 {
	return 0x7ae43737
}

func (*AccountInstallThemeParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountInstallTheme(dark bool, format string, theme InputTheme) (bool, error) {
	responseData, err := c.MakeRequest(&AccountInstallThemeParams{
		Dark:   dark,
		Format: format,
		Theme:  theme,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountInstallTheme")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountInstallWallPaperParams struct {
	Wallpaper InputWallPaper
	Settings  *WallPaperSettings // Settings for the code type to send
}

func (*AccountInstallWallPaperParams) CRC() uint32 {
	return 0xfeed5769
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountInstallWallPaper(wallpaper InputWallPaper, settings *WallPaperSettings) (bool, error) {
	responseData, err := c.MakeRequest(&AccountInstallWallPaperParams{
		Settings:  settings,
		Wallpaper: wallpaper,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountInstallWallPaper")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountRegisterDeviceParams struct {
	NoMuted    bool `tl:"flag:0,encoded_in_bitflags"`
	TokenType  int32
	Token      string
	AppSandbox bool
	Secret     []byte
	OtherUids  []int32
}

func (*AccountRegisterDeviceParams) CRC() uint32 {
	return 0x68976c6f
}

func (*AccountRegisterDeviceParams) FlagIndex() int {
	return 0
}

// Registers a validated phone number in the system.
func (c *Client) AccountRegisterDevice(params *AccountRegisterDeviceParams) (bool, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return false, errors.Wrap(err, "sending AccountRegisterDevice")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountReportPeerParams struct {
	Peer   InputPeer
	Reason ReportReason
}

func (*AccountReportPeerParams) CRC() uint32 {
	return 0xae189d5f
}

// Registers a validated phone number in the system.
func (c *Client) AccountReportPeer(peer InputPeer, reason ReportReason) (bool, error) {
	responseData, err := c.MakeRequest(&AccountReportPeerParams{
		Peer:   peer,
		Reason: reason,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountReportPeer")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountResendPasswordEmailParams struct{}

func (*AccountResendPasswordEmailParams) CRC() uint32 {
	return 0x7a7f2a15
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountResendPasswordEmail() (bool, error) {
	responseData, err := c.MakeRequest(&AccountResendPasswordEmailParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountResendPasswordEmail")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountResetAuthorizationParams struct {
	Hash int64
}

func (*AccountResetAuthorizationParams) CRC() uint32 {
	return 0xdf77f3bc
}

// Registers a validated phone number in the system.
func (c *Client) AccountResetAuthorization(hash int64) (bool, error) {
	responseData, err := c.MakeRequest(&AccountResetAuthorizationParams{Hash: hash})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountResetAuthorization")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountResetNotifySettingsParams struct{}

func (*AccountResetNotifySettingsParams) CRC() uint32 {
	return 0xdb7e1747
}

// Registers a validated phone number in the system.
func (c *Client) AccountResetNotifySettings() (bool, error) {
	responseData, err := c.MakeRequest(&AccountResetNotifySettingsParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountResetNotifySettings")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountResetWallPapersParams struct{}

func (*AccountResetWallPapersParams) CRC() uint32 {
	return 0xbb3b9804
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountResetWallPapers() (bool, error) {
	responseData, err := c.MakeRequest(&AccountResetWallPapersParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountResetWallPapers")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountResetWebAuthorizationParams struct {
	Hash int64
}

func (*AccountResetWebAuthorizationParams) CRC() uint32 {
	return 0x2d01b9ef
}

// Registers a validated phone number in the system.
func (c *Client) AccountResetWebAuthorization(hash int64) (bool, error) {
	responseData, err := c.MakeRequest(&AccountResetWebAuthorizationParams{Hash: hash})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountResetWebAuthorization")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountResetWebAuthorizationsParams struct{}

func (*AccountResetWebAuthorizationsParams) CRC() uint32 {
	return 0x682d2594
}

// Registers a validated phone number in the system.
func (c *Client) AccountResetWebAuthorizations() (bool, error) {
	responseData, err := c.MakeRequest(&AccountResetWebAuthorizationsParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountResetWebAuthorizations")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSaveAutoDownloadSettingsParams struct {
	Low      bool                  `tl:"flag:0,encoded_in_bitflags"`
	High     bool                  `tl:"flag:1,encoded_in_bitflags"`
	Settings *AutoDownloadSettings // Settings for the code type to send
}

func (*AccountSaveAutoDownloadSettingsParams) CRC() uint32 {
	return 0x76f36233
}

func (*AccountSaveAutoDownloadSettingsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSaveAutoDownloadSettings(low, high bool, settings *AutoDownloadSettings) (bool, error) {
	responseData, err := c.MakeRequest(&AccountSaveAutoDownloadSettingsParams{
		High:     high,
		Low:      low,
		Settings: settings,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountSaveAutoDownloadSettings")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountSaveSecureValueParams struct {
	Value          *InputSecureValue
	SecureSecretID int64
}

func (*AccountSaveSecureValueParams) CRC() uint32 {
	return 0x899fe31d
}

// Registers a validated phone number in the system.
func (c *Client) AccountSaveSecureValue(value *InputSecureValue, secureSecretID int64) (*SecureValue, error) {
	responseData, err := c.MakeRequest(&AccountSaveSecureValueParams{
		SecureSecretID: secureSecretID,
		Value:          value,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountSaveSecureValue")
	}

	resp, ok := responseData.(*SecureValue)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSaveThemeParams struct {
	Theme  InputTheme
	Unsave bool
}

func (*AccountSaveThemeParams) CRC() uint32 {
	return 0xf257106c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSaveTheme(theme InputTheme, unsave bool) (bool, error) {
	responseData, err := c.MakeRequest(&AccountSaveThemeParams{
		Theme:  theme,
		Unsave: unsave,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountSaveTheme")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSaveWallPaperParams struct {
	Wallpaper InputWallPaper
	Unsave    bool
	Settings  *WallPaperSettings // Settings for the code type to send
}

func (*AccountSaveWallPaperParams) CRC() uint32 {
	return 0x6c5a5b37
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSaveWallPaper(wallpaper InputWallPaper, unsave bool, settings *WallPaperSettings) (bool, error) {
	responseData, err := c.MakeRequest(&AccountSaveWallPaperParams{
		Settings:  settings,
		Unsave:    unsave,
		Wallpaper: wallpaper,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountSaveWallPaper")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountSendChangePhoneCodeParams struct {
	PhoneNumber string        // Phone number in international format
	Settings    *CodeSettings // Settings for the code type to send
}

func (*AccountSendChangePhoneCodeParams) CRC() uint32 {
	return 0x82574ae5
}

// Registers a validated phone number in the system.
func (c *Client) AccountSendChangePhoneCode(phoneNumber string, settings *CodeSettings) (*AuthSentCode, error) {
	responseData, err := c.MakeRequest(&AccountSendChangePhoneCodeParams{
		PhoneNumber: phoneNumber,
		Settings:    settings,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountSendChangePhoneCode")
	}

	resp, ok := responseData.(*AuthSentCode)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountSendConfirmPhoneCodeParams struct {
	Hash     string
	Settings *CodeSettings // Settings for the code type to send
}

func (*AccountSendConfirmPhoneCodeParams) CRC() uint32 {
	return 0x1b3faa88
}

// Registers a validated phone number in the system.
func (c *Client) AccountSendConfirmPhoneCode(hash string, settings *CodeSettings) (*AuthSentCode, error) {
	responseData, err := c.MakeRequest(&AccountSendConfirmPhoneCodeParams{
		Hash:     hash,
		Settings: settings,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountSendConfirmPhoneCode")
	}

	resp, ok := responseData.(*AuthSentCode)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSendVerifyEmailCodeParams struct {
	Email string
}

func (*AccountSendVerifyEmailCodeParams) CRC() uint32 {
	return 0x7011509f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSendVerifyEmailCode(email string) (*AccountSentEmailCode, error) {
	responseData, err := c.MakeRequest(&AccountSendVerifyEmailCodeParams{Email: email})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountSendVerifyEmailCode")
	}

	resp, ok := responseData.(*AccountSentEmailCode)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSendVerifyPhoneCodeParams struct {
	PhoneNumber string        // Phone number in international format
	Settings    *CodeSettings // Settings for the code type to send
}

func (*AccountSendVerifyPhoneCodeParams) CRC() uint32 {
	return 0xa5a356f9
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSendVerifyPhoneCode(phoneNumber string, settings *CodeSettings) (*AuthSentCode, error) {
	responseData, err := c.MakeRequest(&AccountSendVerifyPhoneCodeParams{
		PhoneNumber: phoneNumber,
		Settings:    settings,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountSendVerifyPhoneCode")
	}

	resp, ok := responseData.(*AuthSentCode)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountSetAccountTtlParams struct {
	Ttl *AccountDaysTtl
}

func (*AccountSetAccountTtlParams) CRC() uint32 {
	return 0x2442485e
}

// Registers a validated phone number in the system.
func (c *Client) AccountSetAccountTtl(ttl *AccountDaysTtl) (bool, error) {
	responseData, err := c.MakeRequest(&AccountSetAccountTtlParams{Ttl: ttl})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountSetAccountTtl")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSetContactSignUpNotificationParams struct {
	Silent bool
}

func (*AccountSetContactSignUpNotificationParams) CRC() uint32 {
	return 0xcff43f61
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSetContactSignUpNotification(silent bool) (bool, error) {
	responseData, err := c.MakeRequest(&AccountSetContactSignUpNotificationParams{Silent: silent})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountSetContactSignUpNotification")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSetContentSettingsParams struct {
	SensitiveEnabled bool `tl:"flag:0,encoded_in_bitflags"`
}

func (*AccountSetContentSettingsParams) CRC() uint32 {
	return 0xb574b16b
}

func (*AccountSetContentSettingsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSetContentSettings(sensitiveEnabled bool) (bool, error) {
	responseData, err := c.MakeRequest(&AccountSetContentSettingsParams{SensitiveEnabled: sensitiveEnabled})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountSetContentSettings")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountSetGlobalPrivacySettingsParams struct {
	Settings *GlobalPrivacySettings // Settings for the code type to send
}

func (*AccountSetGlobalPrivacySettingsParams) CRC() uint32 {
	return 0x1edaaac2
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountSetGlobalPrivacySettings(settings *GlobalPrivacySettings) (*GlobalPrivacySettings, error) {
	responseData, err := c.MakeRequest(&AccountSetGlobalPrivacySettingsParams{Settings: settings})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountSetGlobalPrivacySettings")
	}

	resp, ok := responseData.(*GlobalPrivacySettings)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountSetPrivacyParams struct {
	Key   InputPrivacyKey
	Rules []InputPrivacyRule
}

func (*AccountSetPrivacyParams) CRC() uint32 {
	return 0xc9f81ce8
}

// Registers a validated phone number in the system.
func (c *Client) AccountSetPrivacy(key InputPrivacyKey, rules []InputPrivacyRule) (*AccountPrivacyRules, error) {
	responseData, err := c.MakeRequest(&AccountSetPrivacyParams{
		Key:   key,
		Rules: rules,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountSetPrivacy")
	}

	resp, ok := responseData.(*AccountPrivacyRules)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountUnregisterDeviceParams struct {
	TokenType int32
	Token     string
	OtherUids []int32
}

func (*AccountUnregisterDeviceParams) CRC() uint32 {
	return 0x3076c4bf
}

// Registers a validated phone number in the system.
func (c *Client) AccountUnregisterDevice(tokenType int32, token string, otherUids []int32) (bool, error) {
	responseData, err := c.MakeRequest(&AccountUnregisterDeviceParams{
		OtherUids: otherUids,
		Token:     token,
		TokenType: tokenType,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountUnregisterDevice")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountUpdateDeviceLockedParams struct {
	Period int32
}

func (*AccountUpdateDeviceLockedParams) CRC() uint32 {
	return 0x38df3532
}

// Registers a validated phone number in the system.
func (c *Client) AccountUpdateDeviceLocked(period int32) (bool, error) {
	responseData, err := c.MakeRequest(&AccountUpdateDeviceLockedParams{Period: period})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountUpdateDeviceLocked")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountUpdateNotifySettingsParams struct {
	Peer     InputNotifyPeer
	Settings *InputPeerNotifySettings // Settings for the code type to send
}

func (*AccountUpdateNotifySettingsParams) CRC() uint32 {
	return 0x84be5b93
}

// Registers a validated phone number in the system.
func (c *Client) AccountUpdateNotifySettings(peer InputNotifyPeer, settings *InputPeerNotifySettings) (bool, error) {
	responseData, err := c.MakeRequest(&AccountUpdateNotifySettingsParams{
		Peer:     peer,
		Settings: settings,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountUpdateNotifySettings")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountUpdatePasswordSettingsParams struct {
	Password    InputCheckPasswordSRP
	NewSettings *AccountPasswordInputSettings
}

func (*AccountUpdatePasswordSettingsParams) CRC() uint32 {
	return 0xa59b102f
}

// Registers a validated phone number in the system.
func (c *Client) AccountUpdatePasswordSettings(password InputCheckPasswordSRP, newSettings *AccountPasswordInputSettings) (bool, error) {
	responseData, err := c.MakeRequest(&AccountUpdatePasswordSettingsParams{
		NewSettings: newSettings,
		Password:    password,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountUpdatePasswordSettings")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountUpdateProfileParams struct {
	FirstName string `tl:"flag:0"` // New user first name
	LastName  string `tl:"flag:1"` // New user last name
	About     string `tl:"flag:2"`
}

func (*AccountUpdateProfileParams) CRC() uint32 {
	return 0x78515775
}

func (*AccountUpdateProfileParams) FlagIndex() int {
	return 0
}

// Registers a validated phone number in the system.
func (c *Client) AccountUpdateProfile(firstName, lastName, about string) (User, error) {
	responseData, err := c.MakeRequest(&AccountUpdateProfileParams{
		About:     about,
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountUpdateProfile")
	}

	resp, ok := responseData.(User)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountUpdateStatusParams struct {
	Offline bool
}

func (*AccountUpdateStatusParams) CRC() uint32 {
	return 0x6628562c
}

// Registers a validated phone number in the system.
func (c *Client) AccountUpdateStatus(offline bool) (bool, error) {
	responseData, err := c.MakeRequest(&AccountUpdateStatusParams{Offline: offline})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountUpdateStatus")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountUpdateThemeParams struct {
	Format   string
	Theme    InputTheme
	Slug     string              `tl:"flag:0"`
	Title    string              `tl:"flag:1"`
	Document InputDocument       `tl:"flag:2"`
	Settings *InputThemeSettings `tl:"flag:3"` // Settings for the code type to send
}

func (*AccountUpdateThemeParams) CRC() uint32 {
	return 0x5cb367d5
}

func (*AccountUpdateThemeParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountUpdateTheme(params *AccountUpdateThemeParams) (*Theme, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountUpdateTheme")
	}

	resp, ok := responseData.(*Theme)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AccountUpdateUsernameParams struct {
	Username string
}

func (*AccountUpdateUsernameParams) CRC() uint32 {
	return 0x3e0bdd7c
}

// Registers a validated phone number in the system.
func (c *Client) AccountUpdateUsername(username string) (User, error) {
	responseData, err := c.MakeRequest(&AccountUpdateUsernameParams{Username: username})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountUpdateUsername")
	}

	resp, ok := responseData.(User)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountUploadThemeParams struct {
	File     InputFile
	Thumb    InputFile `tl:"flag:0"`
	FileName string
	MimeType string
}

func (*AccountUploadThemeParams) CRC() uint32 {
	return 0x1c3db333
}

func (*AccountUploadThemeParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountUploadTheme(file, thumb InputFile, fileName, mimeType string) (Document, error) {
	responseData, err := c.MakeRequest(&AccountUploadThemeParams{
		File:     file,
		FileName: fileName,
		MimeType: mimeType,
		Thumb:    thumb,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountUploadTheme")
	}

	resp, ok := responseData.(Document)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountUploadWallPaperParams struct {
	File     InputFile
	MimeType string
	Settings *WallPaperSettings // Settings for the code type to send
}

func (*AccountUploadWallPaperParams) CRC() uint32 {
	return 0xdd853661
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountUploadWallPaper(file InputFile, mimeType string, settings *WallPaperSettings) (WallPaper, error) {
	responseData, err := c.MakeRequest(&AccountUploadWallPaperParams{
		File:     file,
		MimeType: mimeType,
		Settings: settings,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AccountUploadWallPaper")
	}

	resp, ok := responseData.(WallPaper)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountVerifyEmailParams struct {
	Email string
	Code  string
}

func (*AccountVerifyEmailParams) CRC() uint32 {
	return 0xecba39db
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountVerifyEmail(email, code string) (bool, error) {
	responseData, err := c.MakeRequest(&AccountVerifyEmailParams{
		Code:  code,
		Email: email,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountVerifyEmail")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type AccountVerifyPhoneParams struct {
	PhoneNumber   string // Phone number in international format
	PhoneCodeHash string // SMS-message ID
	PhoneCode     string
}

func (*AccountVerifyPhoneParams) CRC() uint32 {
	return 0x4dd3a7f6
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) AccountVerifyPhone(phoneNumber, phoneCodeHash, phoneCode string) (bool, error) {
	responseData, err := c.MakeRequest(&AccountVerifyPhoneParams{
		PhoneCode:     phoneCode,
		PhoneCodeHash: phoneCodeHash,
		PhoneNumber:   phoneNumber,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountVerifyPhone")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthAcceptLoginTokenParams struct {
	Token []byte
}

func (*AuthAcceptLoginTokenParams) CRC() uint32 {
	return 0xe894ad4d
}

// Registers a validated phone number in the system.
func (c *Client) AuthAcceptLoginToken(token []byte) (*Authorization, error) {
	responseData, err := c.MakeRequest(&AuthAcceptLoginTokenParams{Token: token})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthAcceptLoginToken")
	}

	resp, ok := responseData.(*Authorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthBindTempAuthKeyParams struct {
	PermAuthKeyID    int64
	Nonce            int64
	ExpiresAt        int32
	EncryptedMessage []byte
}

func (*AuthBindTempAuthKeyParams) CRC() uint32 {
	return 0xcdd42a05
}

// Registers a validated phone number in the system.
func (c *Client) AuthBindTempAuthKey(permAuthKeyID, nonce int64, expiresAt int32, encryptedMessage []byte) (bool, error) {
	responseData, err := c.MakeRequest(&AuthBindTempAuthKeyParams{
		EncryptedMessage: encryptedMessage,
		ExpiresAt:        expiresAt,
		Nonce:            nonce,
		PermAuthKeyID:    permAuthKeyID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AuthBindTempAuthKey")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthCancelCodeParams struct {
	PhoneNumber   string // Phone number in international format
	PhoneCodeHash string // SMS-message ID
}

func (*AuthCancelCodeParams) CRC() uint32 {
	return 0x1f040578
}

// Registers a validated phone number in the system.
func (c *Client) AuthCancelCode(phoneNumber, phoneCodeHash string) (bool, error) {
	responseData, err := c.MakeRequest(&AuthCancelCodeParams{
		PhoneCodeHash: phoneCodeHash,
		PhoneNumber:   phoneNumber,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending AuthCancelCode")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthCheckPasswordParams struct {
	Password InputCheckPasswordSRP
}

func (*AuthCheckPasswordParams) CRC() uint32 {
	return 0xd18b4d16
}

// Registers a validated phone number in the system.
func (c *Client) AuthCheckPassword(password InputCheckPasswordSRP) (AuthAuthorization, error) {
	responseData, err := c.MakeRequest(&AuthCheckPasswordParams{Password: password})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthCheckPassword")
	}

	resp, ok := responseData.(AuthAuthorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthDropTempAuthKeysParams struct {
	ExceptAuthKeys []int64
}

func (*AuthDropTempAuthKeysParams) CRC() uint32 {
	return 0x8e48a188
}

// Registers a validated phone number in the system.
func (c *Client) AuthDropTempAuthKeys(exceptAuthKeys []int64) (bool, error) {
	responseData, err := c.MakeRequest(&AuthDropTempAuthKeysParams{ExceptAuthKeys: exceptAuthKeys})
	if err != nil {
		return false, errors.Wrap(err, "sending AuthDropTempAuthKeys")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthExportAuthorizationParams struct {
	DcID int32
}

func (*AuthExportAuthorizationParams) CRC() uint32 {
	return 0xe5bfffcd
}

// Registers a validated phone number in the system.
func (c *Client) AuthExportAuthorization(dcID int32) (*AuthExportedAuthorization, error) {
	responseData, err := c.MakeRequest(&AuthExportAuthorizationParams{DcID: dcID})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthExportAuthorization")
	}

	resp, ok := responseData.(*AuthExportedAuthorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthExportLoginTokenParams struct {
	APIID     int32  // Application identifier (see App configuration)
	APIHash   string // Application secret hash (see App configuration)
	ExceptIds []int32
}

func (*AuthExportLoginTokenParams) CRC() uint32 {
	return 0xb1b41517
}

// Registers a validated phone number in the system.
func (c *Client) AuthExportLoginToken(apiID int32, apiHash string, exceptIds []int32) (interface{}, error) {
	responseData, err := c.MakeRequest(&AuthExportLoginTokenParams{
		APIHash:   apiHash,
		APIID:     apiID,
		ExceptIds: exceptIds,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthExportLoginToken")
	}

	return responseData, nil
}

// Registers a validated phone number in the system.
type AuthImportAuthorizationParams struct {
	ID    int32
	Bytes []byte
}

func (*AuthImportAuthorizationParams) CRC() uint32 {
	return 0xe3ef9613
}

// Registers a validated phone number in the system.
func (c *Client) AuthImportAuthorization(id int32, bytes []byte) (AuthAuthorization, error) {
	responseData, err := c.MakeRequest(&AuthImportAuthorizationParams{
		Bytes: bytes,
		ID:    id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthImportAuthorization")
	}

	resp, ok := responseData.(AuthAuthorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthImportBotAuthorizationParams struct {
	Flags        int32  // Flags, see TL conditional fields
	APIID        int32  // Application identifier (see App configuration)
	APIHash      string // Application secret hash (see App configuration)
	BotAuthToken string
}

func (*AuthImportBotAuthorizationParams) CRC() uint32 {
	return 0x67a3ff2c
}

// Registers a validated phone number in the system.
func (c *Client) AuthImportBotAuthorization(flags, apiID int32, apiHash, botAuthToken string) (AuthAuthorization, error) {
	responseData, err := c.MakeRequest(&AuthImportBotAuthorizationParams{
		APIHash:      apiHash,
		APIID:        apiID,
		BotAuthToken: botAuthToken,
		Flags:        flags,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthImportBotAuthorization")
	}

	resp, ok := responseData.(AuthAuthorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthImportLoginTokenParams struct {
	Token []byte
}

func (*AuthImportLoginTokenParams) CRC() uint32 {
	return 0x95ac5ce4
}

// Registers a validated phone number in the system.
func (c *Client) AuthImportLoginToken(token []byte) (interface{}, error) {
	responseData, err := c.MakeRequest(&AuthImportLoginTokenParams{Token: token})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthImportLoginToken")
	}
	log.Println("FINISHED")

	return responseData, nil
}

// Registers a validated phone number in the system.
type AuthLogOutParams struct{}

func (*AuthLogOutParams) CRC() uint32 {
	return 0x5717da40
}

// Registers a validated phone number in the system.
func (c *Client) AuthLogOut() (bool, error) {
	responseData, err := c.MakeRequest(&AuthLogOutParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AuthLogOut")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthRecoverPasswordParams struct {
	Code string
}

func (*AuthRecoverPasswordParams) CRC() uint32 {
	return 0x4ea56e92
}

// Registers a validated phone number in the system.
func (c *Client) AuthRecoverPassword(code string) (AuthAuthorization, error) {
	responseData, err := c.MakeRequest(&AuthRecoverPasswordParams{Code: code})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthRecoverPassword")
	}

	resp, ok := responseData.(AuthAuthorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthRequestPasswordRecoveryParams struct{}

func (*AuthRequestPasswordRecoveryParams) CRC() uint32 {
	return 0xd897bc66
}

// Registers a validated phone number in the system.
func (c *Client) AuthRequestPasswordRecovery() (*AuthPasswordRecovery, error) {
	responseData, err := c.MakeRequest(&AuthRequestPasswordRecoveryParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthRequestPasswordRecovery")
	}

	resp, ok := responseData.(*AuthPasswordRecovery)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthResendCodeParams struct {
	PhoneNumber   string // Phone number in international format
	PhoneCodeHash string // SMS-message ID
}

func (*AuthResendCodeParams) CRC() uint32 {
	return 0x3ef1a9bf
}

// Registers a validated phone number in the system.
func (c *Client) AuthResendCode(phoneNumber, phoneCodeHash string) (*AuthSentCode, error) {
	responseData, err := c.MakeRequest(&AuthResendCodeParams{
		PhoneCodeHash: phoneCodeHash,
		PhoneNumber:   phoneNumber,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthResendCode")
	}

	resp, ok := responseData.(*AuthSentCode)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthResetAuthorizationsParams struct{}

func (*AuthResetAuthorizationsParams) CRC() uint32 {
	return 0x9fab0d1a
}

// Registers a validated phone number in the system.
func (c *Client) AuthResetAuthorizations() (bool, error) {
	responseData, err := c.MakeRequest(&AuthResetAuthorizationsParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending AuthResetAuthorizations")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Send the verification code for login
type AuthSendCodeParams struct {
	PhoneNumber string        // Phone number in international format
	APIID       int32         // Application identifier (see App configuration)
	APIHash     string        // Application secret hash (see App configuration)
	Settings    *CodeSettings // Settings for the code type to send
}

func (*AuthSendCodeParams) CRC() uint32 {
	return 0xa677244f
}

// Send the verification code for login
func (c *Client) AuthSendCode(phoneNumber string, apiID int32, apiHash string, settings *CodeSettings) (*AuthSentCode, error) {
	responseData, err := c.MakeRequest(&AuthSendCodeParams{
		APIHash:     apiHash,
		APIID:       apiID,
		PhoneNumber: phoneNumber,
		Settings:    settings,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthSendCode")
	}

	resp, ok := responseData.(*AuthSentCode)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthSignInParams struct {
	PhoneNumber   string // Phone number in international format
	PhoneCodeHash string // SMS-message ID
	PhoneCode     string
}

func (*AuthSignInParams) CRC() uint32 {
	return 0xbcd51581
}

// Registers a validated phone number in the system.
func (c *Client) AuthSignIn(phoneNumber, phoneCodeHash, phoneCode string) (AuthAuthorization, error) {
	responseData, err := c.MakeRequest(&AuthSignInParams{
		PhoneCode:     phoneCode,
		PhoneCodeHash: phoneCodeHash,
		PhoneNumber:   phoneNumber,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthSignIn")
	}

	resp, ok := responseData.(AuthAuthorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Registers a validated phone number in the system.
type AuthSignUpParams struct {
	PhoneNumber   string // Phone number in international format
	PhoneCodeHash string // SMS-message ID
	FirstName     string // New user first name
	LastName      string // New user last name
}

func (*AuthSignUpParams) CRC() uint32 {
	return 0x80eee427
}

// Registers a validated phone number in the system.
func (c *Client) AuthSignUp(phoneNumber, phoneCodeHash, firstName, lastName string) (AuthAuthorization, error) {
	responseData, err := c.MakeRequest(&AuthSignUpParams{
		FirstName:     firstName,
		LastName:      lastName,
		PhoneCodeHash: phoneCodeHash,
		PhoneNumber:   phoneNumber,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending AuthSignUp")
	}

	resp, ok := responseData.(AuthAuthorization)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type BotsAnswerWebhookJsonQueryParams struct {
	QueryID int64
	Data    *DataJson
}

func (*BotsAnswerWebhookJsonQueryParams) CRC() uint32 {
	return 0xe6213f4d
}

// Get the participants of a channel
func (c *Client) BotsAnswerWebhookJsonQuery(queryID int64, data *DataJson) (bool, error) {
	responseData, err := c.MakeRequest(&BotsAnswerWebhookJsonQueryParams{
		Data:    data,
		QueryID: queryID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending BotsAnswerWebhookJsonQuery")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type BotsSendCustomRequestParams struct {
	CustomMethod string
	Params       *DataJson // Additional initConnection parameters. For now, only the tz_offset field is supported, for specifying timezone offset in seconds.
}

func (*BotsSendCustomRequestParams) CRC() uint32 {
	return 0xaa2769ed
}

// Get the participants of a channel
func (c *Client) BotsSendCustomRequest(customMethod string, params *DataJson) (*DataJson, error) {
	responseData, err := c.MakeRequest(&BotsSendCustomRequestParams{
		CustomMethod: customMethod,
		Params:       params,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending BotsSendCustomRequest")
	}

	resp, ok := responseData.(*DataJson)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type BotsSetBotCommandsParams struct {
	Commands []*BotCommand
}

func (*BotsSetBotCommandsParams) CRC() uint32 {
	return 0x805d46f6
}

// Get the participants of a channel
func (c *Client) BotsSetBotCommands(commands []*BotCommand) (bool, error) {
	responseData, err := c.MakeRequest(&BotsSetBotCommandsParams{Commands: commands})
	if err != nil {
		return false, errors.Wrap(err, "sending BotsSetBotCommands")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsCheckUsernameParams struct {
	Channel  InputChannel // Channel
	Username string
}

func (*ChannelsCheckUsernameParams) CRC() uint32 {
	return 0x10e6bd2c
}

// Get the participants of a channel
func (c *Client) ChannelsCheckUsername(channel InputChannel, username string) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsCheckUsernameParams{
		Channel:  channel,
		Username: username,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsCheckUsername")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsCreateChannelParams struct {
	Broadcast bool `tl:"flag:0,encoded_in_bitflags"`
	Megagroup bool `tl:"flag:1,encoded_in_bitflags"`
	Title     string
	About     string
	GeoPoint  InputGeoPoint `tl:"flag:2"`
	Address   string        `tl:"flag:2"`
}

func (*ChannelsCreateChannelParams) CRC() uint32 {
	return 0x3d5fb10f
}

func (*ChannelsCreateChannelParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) ChannelsCreateChannel(params *ChannelsCreateChannelParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsCreateChannel")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsDeleteChannelParams struct {
	Channel InputChannel // Channel
}

func (*ChannelsDeleteChannelParams) CRC() uint32 {
	return 0xc0111fe3
}

// Get the participants of a channel
func (c *Client) ChannelsDeleteChannel(channel InputChannel) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsDeleteChannelParams{Channel: channel})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsDeleteChannel")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsDeleteHistoryParams struct {
	Channel InputChannel // Channel
	MaxID   int32
}

func (*ChannelsDeleteHistoryParams) CRC() uint32 {
	return 0xaf369d42
}

// Get the participants of a channel
func (c *Client) ChannelsDeleteHistory(channel InputChannel, maxID int32) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsDeleteHistoryParams{
		Channel: channel,
		MaxID:   maxID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsDeleteHistory")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ChannelsDeleteMessagesParams struct {
	Channel InputChannel
	ID      []int32
}

func (*ChannelsDeleteMessagesParams) CRC() uint32 {
	return 0x84c1fd4e
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ChannelsDeleteMessages(channel InputChannel, id []int32) (*MessagesAffectedMessages, error) {
	responseData, err := c.MakeRequest(&ChannelsDeleteMessagesParams{
		Channel: channel,
		ID:      id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsDeleteMessages")
	}

	resp, ok := responseData.(*MessagesAffectedMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ChannelsDeleteUserHistoryParams struct {
	Channel InputChannel
	UserID  InputUser
}

func (*ChannelsDeleteUserHistoryParams) CRC() uint32 {
	return 0xd10dd71b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ChannelsDeleteUserHistory(channel InputChannel, userID InputUser) (*MessagesAffectedHistory, error) {
	responseData, err := c.MakeRequest(&ChannelsDeleteUserHistoryParams{
		Channel: channel,
		UserID:  userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsDeleteUserHistory")
	}

	resp, ok := responseData.(*MessagesAffectedHistory)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsEditAdminParams struct {
	Channel     InputChannel // Channel
	UserID      InputUser
	AdminRights *ChatAdminRights
	Rank        string
}

func (*ChannelsEditAdminParams) CRC() uint32 {
	return 0xd33c8902
}

// Get the participants of a channel
func (c *Client) ChannelsEditAdmin(channel InputChannel, userID InputUser, adminRights *ChatAdminRights, rank string) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsEditAdminParams{
		AdminRights: adminRights,
		Channel:     channel,
		Rank:        rank,
		UserID:      userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsEditAdmin")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsEditBannedParams struct {
	Channel      InputChannel // Channel
	UserID       InputUser
	BannedRights *ChatBannedRights
}

func (*ChannelsEditBannedParams) CRC() uint32 {
	return 0x72796912
}

// Get the participants of a channel
func (c *Client) ChannelsEditBanned(channel InputChannel, userID InputUser, bannedRights *ChatBannedRights) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsEditBannedParams{
		BannedRights: bannedRights,
		Channel:      channel,
		UserID:       userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsEditBanned")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsEditCreatorParams struct {
	Channel  InputChannel // Channel
	UserID   InputUser
	Password InputCheckPasswordSRP
}

func (*ChannelsEditCreatorParams) CRC() uint32 {
	return 0x8f38cd1f
}

// Get the participants of a channel
func (c *Client) ChannelsEditCreator(channel InputChannel, userID InputUser, password InputCheckPasswordSRP) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsEditCreatorParams{
		Channel:  channel,
		Password: password,
		UserID:   userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsEditCreator")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsEditLocationParams struct {
	Channel  InputChannel // Channel
	GeoPoint InputGeoPoint
	Address  string
}

func (*ChannelsEditLocationParams) CRC() uint32 {
	return 0x58e63f6d
}

// Get the participants of a channel
func (c *Client) ChannelsEditLocation(channel InputChannel, geoPoint InputGeoPoint, address string) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsEditLocationParams{
		Address:  address,
		Channel:  channel,
		GeoPoint: geoPoint,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsEditLocation")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsEditPhotoParams struct {
	Channel InputChannel // Channel
	Photo   InputChatPhoto
}

func (*ChannelsEditPhotoParams) CRC() uint32 {
	return 0xf12e57c9
}

// Get the participants of a channel
func (c *Client) ChannelsEditPhoto(channel InputChannel, photo InputChatPhoto) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsEditPhotoParams{
		Channel: channel,
		Photo:   photo,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsEditPhoto")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsEditTitleParams struct {
	Channel InputChannel // Channel
	Title   string
}

func (*ChannelsEditTitleParams) CRC() uint32 {
	return 0x566decd0
}

// Get the participants of a channel
func (c *Client) ChannelsEditTitle(channel InputChannel, title string) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsEditTitleParams{
		Channel: channel,
		Title:   title,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsEditTitle")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsExportMessageLinkParams struct {
	Grouped bool         `tl:"flag:0,encoded_in_bitflags"`
	Thread  bool         `tl:"flag:1,encoded_in_bitflags"`
	Channel InputChannel // Channel
	ID      int32
}

func (*ChannelsExportMessageLinkParams) CRC() uint32 {
	return 0xe63fadeb
}

func (*ChannelsExportMessageLinkParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) ChannelsExportMessageLink(grouped, thread bool, channel InputChannel, id int32) (*ExportedMessageLink, error) {
	responseData, err := c.MakeRequest(&ChannelsExportMessageLinkParams{
		Channel: channel,
		Grouped: grouped,
		ID:      id,
		Thread:  thread,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsExportMessageLink")
	}

	resp, ok := responseData.(*ExportedMessageLink)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetAdminLogParams struct {
	Channel      InputChannel // Channel
	Q            string
	EventsFilter *ChannelAdminLogEventsFilter `tl:"flag:0"`
	Admins       []InputUser                  `tl:"flag:1"`
	MaxID        int64
	MinID        int64
	Limit        int32
}

func (*ChannelsGetAdminLogParams) CRC() uint32 {
	return 0x33ddf480
}

func (*ChannelsGetAdminLogParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) ChannelsGetAdminLog(params *ChannelsGetAdminLogParams) (*ChannelsAdminLogResults, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetAdminLog")
	}

	resp, ok := responseData.(*ChannelsAdminLogResults)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetAdminedPublicChannelsParams struct {
	ByLocation bool `tl:"flag:0,encoded_in_bitflags"`
	CheckLimit bool `tl:"flag:1,encoded_in_bitflags"`
}

func (*ChannelsGetAdminedPublicChannelsParams) CRC() uint32 {
	return 0xf8b036af
}

func (*ChannelsGetAdminedPublicChannelsParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) ChannelsGetAdminedPublicChannels(byLocation, checkLimit bool) (MessagesChats, error) {
	responseData, err := c.MakeRequest(&ChannelsGetAdminedPublicChannelsParams{
		ByLocation: byLocation,
		CheckLimit: checkLimit,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetAdminedPublicChannels")
	}

	resp, ok := responseData.(MessagesChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetChannelsParams struct {
	ID []InputChannel
}

func (*ChannelsGetChannelsParams) CRC() uint32 {
	return 0xa7f6bbb
}

// Get the participants of a channel
func (c *Client) ChannelsGetChannels(id []InputChannel) (MessagesChats, error) {
	responseData, err := c.MakeRequest(&ChannelsGetChannelsParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetChannels")
	}

	resp, ok := responseData.(MessagesChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetFullChannelParams struct {
	Channel InputChannel // Channel
}

func (*ChannelsGetFullChannelParams) CRC() uint32 {
	return 0x8736a09
}

// Get the participants of a channel
func (c *Client) ChannelsGetFullChannel(channel InputChannel) (*MessagesChatFull, error) {
	responseData, err := c.MakeRequest(&ChannelsGetFullChannelParams{Channel: channel})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetFullChannel")
	}

	resp, ok := responseData.(*MessagesChatFull)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetGroupsForDiscussionParams struct{}

func (*ChannelsGetGroupsForDiscussionParams) CRC() uint32 {
	return 0xf5dad378
}

// Get the participants of a channel
func (c *Client) ChannelsGetGroupsForDiscussion() (MessagesChats, error) {
	responseData, err := c.MakeRequest(&ChannelsGetGroupsForDiscussionParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetGroupsForDiscussion")
	}

	resp, ok := responseData.(MessagesChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetInactiveChannelsParams struct{}

func (*ChannelsGetInactiveChannelsParams) CRC() uint32 {
	return 0x11e831ee
}

// Get the participants of a channel
func (c *Client) ChannelsGetInactiveChannels() (*MessagesInactiveChats, error) {
	responseData, err := c.MakeRequest(&ChannelsGetInactiveChannelsParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetInactiveChannels")
	}

	resp, ok := responseData.(*MessagesInactiveChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetLeftChannelsParams struct {
	Offset int32
}

func (*ChannelsGetLeftChannelsParams) CRC() uint32 {
	return 0x8341ecc0
}

// Get the participants of a channel
func (c *Client) ChannelsGetLeftChannels(offset int32) (MessagesChats, error) {
	responseData, err := c.MakeRequest(&ChannelsGetLeftChannelsParams{Offset: offset})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetLeftChannels")
	}

	resp, ok := responseData.(MessagesChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ChannelsGetMessagesParams struct {
	Channel InputChannel
	ID      []InputMessage
}

func (*ChannelsGetMessagesParams) CRC() uint32 {
	return 0xad8c9a23
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ChannelsGetMessages(channel InputChannel, id []InputMessage) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(&ChannelsGetMessagesParams{
		Channel: channel,
		ID:      id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetMessages")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetParticipantParams struct {
	Channel InputChannel // Channel
	UserID  InputUser
}

func (*ChannelsGetParticipantParams) CRC() uint32 {
	return 0x546dd7a6
}

// Get the participants of a channel
func (c *Client) ChannelsGetParticipant(channel InputChannel, userID InputUser) (*ChannelsChannelParticipant, error) {
	responseData, err := c.MakeRequest(&ChannelsGetParticipantParams{
		Channel: channel,
		UserID:  userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetParticipant")
	}

	resp, ok := responseData.(*ChannelsChannelParticipant)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsGetParticipantsParams struct {
	Channel InputChannel              // Channel
	Filter  ChannelParticipantsFilter // Which participant types to fetch
	Offset  int32
	Limit   int32
	Hash    int32
}

func (*ChannelsGetParticipantsParams) CRC() uint32 {
	return 0x123e05e9
}

// Get the participants of a channel
func (c *Client) ChannelsGetParticipants(channel InputChannel, filter ChannelParticipantsFilter, offset, limit, hash int32) (ChannelsChannelParticipants, error) {
	responseData, err := c.MakeRequest(&ChannelsGetParticipantsParams{
		Channel: channel,
		Filter:  filter,
		Hash:    hash,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsGetParticipants")
	}

	resp, ok := responseData.(ChannelsChannelParticipants)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsInviteToChannelParams struct {
	Channel InputChannel // Channel
	Users   []InputUser
}

func (*ChannelsInviteToChannelParams) CRC() uint32 {
	return 0x199f3a6c
}

// Get the participants of a channel
func (c *Client) ChannelsInviteToChannel(channel InputChannel, users []InputUser) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsInviteToChannelParams{
		Channel: channel,
		Users:   users,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsInviteToChannel")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsJoinChannelParams struct {
	Channel InputChannel // Channel
}

func (*ChannelsJoinChannelParams) CRC() uint32 {
	return 0x24b524c5
}

// Get the participants of a channel
func (c *Client) ChannelsJoinChannel(channel InputChannel) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsJoinChannelParams{Channel: channel})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsJoinChannel")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsLeaveChannelParams struct {
	Channel InputChannel // Channel
}

func (*ChannelsLeaveChannelParams) CRC() uint32 {
	return 0xf836aa95
}

// Get the participants of a channel
func (c *Client) ChannelsLeaveChannel(channel InputChannel) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsLeaveChannelParams{Channel: channel})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsLeaveChannel")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ChannelsReadHistoryParams struct {
	Channel InputChannel
	MaxID   int32
}

func (*ChannelsReadHistoryParams) CRC() uint32 {
	return 0xcc104937
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ChannelsReadHistory(channel InputChannel, maxID int32) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsReadHistoryParams{
		Channel: channel,
		MaxID:   maxID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsReadHistory")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsReadMessageContentsParams struct {
	Channel InputChannel // Channel
	ID      []int32
}

func (*ChannelsReadMessageContentsParams) CRC() uint32 {
	return 0xeab5dc38
}

// Get the participants of a channel
func (c *Client) ChannelsReadMessageContents(channel InputChannel, id []int32) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsReadMessageContentsParams{
		Channel: channel,
		ID:      id,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsReadMessageContents")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ChannelsReportSpamParams struct {
	Channel InputChannel
	UserID  InputUser
	ID      []int32
}

func (*ChannelsReportSpamParams) CRC() uint32 {
	return 0xfe087810
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ChannelsReportSpam(channel InputChannel, userID InputUser, id []int32) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsReportSpamParams{
		Channel: channel,
		ID:      id,
		UserID:  userID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsReportSpam")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsSetDiscussionGroupParams struct {
	Broadcast InputChannel
	Group     InputChannel
}

func (*ChannelsSetDiscussionGroupParams) CRC() uint32 {
	return 0x40582bb2
}

// Get the participants of a channel
func (c *Client) ChannelsSetDiscussionGroup(broadcast, group InputChannel) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsSetDiscussionGroupParams{
		Broadcast: broadcast,
		Group:     group,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsSetDiscussionGroup")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsSetStickersParams struct {
	Channel    InputChannel // Channel
	Stickerset InputStickerSet
}

func (*ChannelsSetStickersParams) CRC() uint32 {
	return 0xea8ca4f9
}

// Get the participants of a channel
func (c *Client) ChannelsSetStickers(channel InputChannel, stickerset InputStickerSet) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsSetStickersParams{
		Channel:    channel,
		Stickerset: stickerset,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsSetStickers")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsTogglePreHistoryHiddenParams struct {
	Channel InputChannel // Channel
	Enabled bool
}

func (*ChannelsTogglePreHistoryHiddenParams) CRC() uint32 {
	return 0xeabbb94c
}

// Get the participants of a channel
func (c *Client) ChannelsTogglePreHistoryHidden(channel InputChannel, enabled bool) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsTogglePreHistoryHiddenParams{
		Channel: channel,
		Enabled: enabled,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsTogglePreHistoryHidden")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsToggleSignaturesParams struct {
	Channel InputChannel // Channel
	Enabled bool
}

func (*ChannelsToggleSignaturesParams) CRC() uint32 {
	return 0x1f69b606
}

// Get the participants of a channel
func (c *Client) ChannelsToggleSignatures(channel InputChannel, enabled bool) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsToggleSignaturesParams{
		Channel: channel,
		Enabled: enabled,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsToggleSignatures")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsToggleSlowModeParams struct {
	Channel InputChannel // Channel
	Seconds int32
}

func (*ChannelsToggleSlowModeParams) CRC() uint32 {
	return 0xedd49ef0
}

// Get the participants of a channel
func (c *Client) ChannelsToggleSlowMode(channel InputChannel, seconds int32) (Updates, error) {
	responseData, err := c.MakeRequest(&ChannelsToggleSlowModeParams{
		Channel: channel,
		Seconds: seconds,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ChannelsToggleSlowMode")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type ChannelsUpdateUsernameParams struct {
	Channel  InputChannel // Channel
	Username string
}

func (*ChannelsUpdateUsernameParams) CRC() uint32 {
	return 0x3514b3de
}

// Get the participants of a channel
func (c *Client) ChannelsUpdateUsername(channel InputChannel, username string) (bool, error) {
	responseData, err := c.MakeRequest(&ChannelsUpdateUsernameParams{
		Channel:  channel,
		Username: username,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ChannelsUpdateUsername")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsAcceptContactParams struct {
	ID InputUser
}

func (*ContactsAcceptContactParams) CRC() uint32 {
	return 0xf831a20f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsAcceptContact(id InputUser) (Updates, error) {
	responseData, err := c.MakeRequest(&ContactsAcceptContactParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsAcceptContact")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsAddContactParams struct {
	AddPhonePrivacyException bool `tl:"flag:0,encoded_in_bitflags"`
	ID                       InputUser
	FirstName                string // New user first name
	LastName                 string // New user last name
	Phone                    string
}

func (*ContactsAddContactParams) CRC() uint32 {
	return 0xe8f463d0
}

func (*ContactsAddContactParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsAddContact(params *ContactsAddContactParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsAddContact")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsBlockParams struct {
	ID InputPeer
}

func (*ContactsBlockParams) CRC() uint32 {
	return 0x68cc1411
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsBlock(id InputPeer) (bool, error) {
	responseData, err := c.MakeRequest(&ContactsBlockParams{ID: id})
	if err != nil {
		return false, errors.Wrap(err, "sending ContactsBlock")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsBlockFromRepliesParams struct {
	DeleteMessage bool  `tl:"flag:0,encoded_in_bitflags"`
	DeleteHistory bool  `tl:"flag:1,encoded_in_bitflags"`
	ReportSpam    bool  `tl:"flag:2,encoded_in_bitflags"`
	MsgID         int32 // Message identifier on which a current query depends
}

func (*ContactsBlockFromRepliesParams) CRC() uint32 {
	return 0x29a8962c
}

func (*ContactsBlockFromRepliesParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsBlockFromReplies(deleteMessage, deleteHistory, reportSpam bool, msgID int32) (Updates, error) {
	responseData, err := c.MakeRequest(&ContactsBlockFromRepliesParams{
		DeleteHistory: deleteHistory,
		DeleteMessage: deleteMessage,
		MsgID:         msgID,
		ReportSpam:    reportSpam,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsBlockFromReplies")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsDeleteByPhonesParams struct {
	Phones []string
}

func (*ContactsDeleteByPhonesParams) CRC() uint32 {
	return 0x1013fd9e
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsDeleteByPhones(phones []string) (bool, error) {
	responseData, err := c.MakeRequest(&ContactsDeleteByPhonesParams{Phones: phones})
	if err != nil {
		return false, errors.Wrap(err, "sending ContactsDeleteByPhones")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsDeleteContactsParams struct {
	ID []InputUser
}

func (*ContactsDeleteContactsParams) CRC() uint32 {
	return 0x96a0e00
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsDeleteContacts(id []InputUser) (Updates, error) {
	responseData, err := c.MakeRequest(&ContactsDeleteContactsParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsDeleteContacts")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsGetBlockedParams struct {
	Offset int32
	Limit  int32
}

func (*ContactsGetBlockedParams) CRC() uint32 {
	return 0xf57c350f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsGetBlocked(offset, limit int32) (ContactsBlocked, error) {
	responseData, err := c.MakeRequest(&ContactsGetBlockedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsGetBlocked")
	}

	resp, ok := responseData.(ContactsBlocked)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsGetContactIDsParams struct {
	Hash int32
}

func (*ContactsGetContactIDsParams) CRC() uint32 {
	return 0x2caa4a42
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsGetContactIDs(hash int32) ([]int32, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&ContactsGetContactIDsParams{Hash: hash}, reflect.TypeOf([]int32{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsGetContactIDs")
	}

	resp, ok := responseData.([]int32)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsGetContactsParams struct {
	Hash int32
}

func (*ContactsGetContactsParams) CRC() uint32 {
	return 0xc023849f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsGetContacts(hash int32) (ContactsContacts, error) {
	responseData, err := c.MakeRequest(&ContactsGetContactsParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsGetContacts")
	}

	resp, ok := responseData.(ContactsContacts)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsGetLocatedParams struct {
	Background  bool `tl:"flag:1,encoded_in_bitflags"`
	GeoPoint    InputGeoPoint
	SelfExpires int32 `tl:"flag:0"`
}

func (*ContactsGetLocatedParams) CRC() uint32 {
	return 0xd348bc44
}

func (*ContactsGetLocatedParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsGetLocated(background bool, geoPoint InputGeoPoint, selfExpires int32) (Updates, error) {
	responseData, err := c.MakeRequest(&ContactsGetLocatedParams{
		Background:  background,
		GeoPoint:    geoPoint,
		SelfExpires: selfExpires,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsGetLocated")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsGetSavedParams struct{}

func (*ContactsGetSavedParams) CRC() uint32 {
	return 0x82f1e39f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsGetSaved() ([]*SavedPhoneContact, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&ContactsGetSavedParams{}, reflect.TypeOf([]*SavedPhoneContact{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsGetSaved")
	}

	resp, ok := responseData.([]*SavedPhoneContact)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsGetStatusesParams struct{}

func (*ContactsGetStatusesParams) CRC() uint32 {
	return 0xc4a353ee
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsGetStatuses() ([]*ContactStatus, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&ContactsGetStatusesParams{}, reflect.TypeOf([]*ContactStatus{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsGetStatuses")
	}

	resp, ok := responseData.([]*ContactStatus)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsGetTopPeersParams struct {
	Correspondents bool `tl:"flag:0,encoded_in_bitflags"`
	BotsPm         bool `tl:"flag:1,encoded_in_bitflags"`
	BotsInline     bool `tl:"flag:2,encoded_in_bitflags"`
	PhoneCalls     bool `tl:"flag:3,encoded_in_bitflags"`
	ForwardUsers   bool `tl:"flag:4,encoded_in_bitflags"`
	ForwardChats   bool `tl:"flag:5,encoded_in_bitflags"`
	Groups         bool `tl:"flag:10,encoded_in_bitflags"`
	Channels       bool `tl:"flag:15,encoded_in_bitflags"`
	Offset         int32
	Limit          int32
	Hash           int32
}

func (*ContactsGetTopPeersParams) CRC() uint32 {
	return 0xd4982db5
}

func (*ContactsGetTopPeersParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsGetTopPeers(params *ContactsGetTopPeersParams) (ContactsTopPeers, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsGetTopPeers")
	}

	resp, ok := responseData.(ContactsTopPeers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsImportContactsParams struct {
	Contacts []*InputPhoneContact
}

func (*ContactsImportContactsParams) CRC() uint32 {
	return 0x2c800be5
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsImportContacts(contacts []*InputPhoneContact) (*ContactsImportedContacts, error) {
	responseData, err := c.MakeRequest(&ContactsImportContactsParams{Contacts: contacts})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsImportContacts")
	}

	resp, ok := responseData.(*ContactsImportedContacts)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsResetSavedParams struct{}

func (*ContactsResetSavedParams) CRC() uint32 {
	return 0x879537f1
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsResetSaved() (bool, error) {
	responseData, err := c.MakeRequest(&ContactsResetSavedParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending ContactsResetSaved")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsResetTopPeerRatingParams struct {
	Category TopPeerCategory
	Peer     InputPeer
}

func (*ContactsResetTopPeerRatingParams) CRC() uint32 {
	return 0x1ae373ac
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsResetTopPeerRating(category TopPeerCategory, peer InputPeer) (bool, error) {
	responseData, err := c.MakeRequest(&ContactsResetTopPeerRatingParams{
		Category: category,
		Peer:     peer,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending ContactsResetTopPeerRating")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsResolveUsernameParams struct {
	Username string
}

func (*ContactsResolveUsernameParams) CRC() uint32 {
	return 0xf93ccba3
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsResolveUsername(username string) (*ContactsResolvedPeer, error) {
	responseData, err := c.MakeRequest(&ContactsResolveUsernameParams{Username: username})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsResolveUsername")
	}

	resp, ok := responseData.(*ContactsResolvedPeer)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsSearchParams struct {
	Q     string
	Limit int32
}

func (*ContactsSearchParams) CRC() uint32 {
	return 0x11f812d8
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsSearch(q string, limit int32) (*ContactsFound, error) {
	responseData, err := c.MakeRequest(&ContactsSearchParams{
		Limit: limit,
		Q:     q,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending ContactsSearch")
	}

	resp, ok := responseData.(*ContactsFound)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsToggleTopPeersParams struct {
	Enabled bool
}

func (*ContactsToggleTopPeersParams) CRC() uint32 {
	return 0x8514bdda
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsToggleTopPeers(enabled bool) (bool, error) {
	responseData, err := c.MakeRequest(&ContactsToggleTopPeersParams{Enabled: enabled})
	if err != nil {
		return false, errors.Wrap(err, "sending ContactsToggleTopPeers")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type ContactsUnblockParams struct {
	ID InputPeer
}

func (*ContactsUnblockParams) CRC() uint32 {
	return 0xbea65d50
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) ContactsUnblock(id InputPeer) (bool, error) {
	responseData, err := c.MakeRequest(&ContactsUnblockParams{ID: id})
	if err != nil {
		return false, errors.Wrap(err, "sending ContactsUnblock")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type FoldersDeleteFolderParams struct {
	FolderID int32
}

func (*FoldersDeleteFolderParams) CRC() uint32 {
	return 0x1c295881
}

// Get the participants of a channel
func (c *Client) FoldersDeleteFolder(folderID int32) (Updates, error) {
	responseData, err := c.MakeRequest(&FoldersDeleteFolderParams{FolderID: folderID})
	if err != nil {
		return nil, errors.Wrap(err, "sending FoldersDeleteFolder")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type FoldersEditPeerFoldersParams struct {
	FolderPeers []*InputFolderPeer
}

func (*FoldersEditPeerFoldersParams) CRC() uint32 {
	return 0x6847d0ab
}

// Get the participants of a channel
func (c *Client) FoldersEditPeerFolders(folderPeers []*InputFolderPeer) (Updates, error) {
	responseData, err := c.MakeRequest(&FoldersEditPeerFoldersParams{FolderPeers: folderPeers})
	if err != nil {
		return nil, errors.Wrap(err, "sending FoldersEditPeerFolders")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpAcceptTermsOfServiceParams struct {
	ID *DataJson
}

func (*HelpAcceptTermsOfServiceParams) CRC() uint32 {
	return 0xee72f79a
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpAcceptTermsOfService(id *DataJson) (bool, error) {
	responseData, err := c.MakeRequest(&HelpAcceptTermsOfServiceParams{ID: id})
	if err != nil {
		return false, errors.Wrap(err, "sending HelpAcceptTermsOfService")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpDismissSuggestionParams struct {
	Suggestion string
}

func (*HelpDismissSuggestionParams) CRC() uint32 {
	return 0x77fa99f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpDismissSuggestion(suggestion string) (bool, error) {
	responseData, err := c.MakeRequest(&HelpDismissSuggestionParams{Suggestion: suggestion})
	if err != nil {
		return false, errors.Wrap(err, "sending HelpDismissSuggestion")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpEditUserInfoParams struct {
	UserID   InputUser
	Message  string
	Entities []MessageEntity
}

func (*HelpEditUserInfoParams) CRC() uint32 {
	return 0x66b91b70
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpEditUserInfo(userID InputUser, message string, entities []MessageEntity) (HelpUserInfo, error) {
	responseData, err := c.MakeRequest(&HelpEditUserInfoParams{
		Entities: entities,
		Message:  message,
		UserID:   userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpEditUserInfo")
	}

	resp, ok := responseData.(HelpUserInfo)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetAppChangelogParams struct {
	PrevAppVersion string
}

func (*HelpGetAppChangelogParams) CRC() uint32 {
	return 0x9010ef6f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetAppChangelog(prevAppVersion string) (Updates, error) {
	responseData, err := c.MakeRequest(&HelpGetAppChangelogParams{PrevAppVersion: prevAppVersion})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetAppChangelog")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetAppConfigParams struct{}

func (*HelpGetAppConfigParams) CRC() uint32 {
	return 0x98914110
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetAppConfig() (JsonValue, error) {
	responseData, err := c.MakeRequest(&HelpGetAppConfigParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetAppConfig")
	}

	resp, ok := responseData.(JsonValue)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetAppUpdateParams struct {
	Source string
}

func (*HelpGetAppUpdateParams) CRC() uint32 {
	return 0x522d5a7d
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetAppUpdate(source string) (HelpAppUpdate, error) {
	responseData, err := c.MakeRequest(&HelpGetAppUpdateParams{Source: source})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetAppUpdate")
	}

	resp, ok := responseData.(HelpAppUpdate)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetCdnConfigParams struct{}

func (*HelpGetCdnConfigParams) CRC() uint32 {
	return 0x52029342
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetCdnConfig() (*CdnConfig, error) {
	responseData, err := c.MakeRequest(&HelpGetCdnConfigParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetCdnConfig")
	}

	resp, ok := responseData.(*CdnConfig)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetConfigParams struct{}

func (*HelpGetConfigParams) CRC() uint32 {
	return 0xc4f9186b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetConfig() (*Config, error) {
	responseData, err := c.MakeRequest(&HelpGetConfigParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetConfig")
	}

	resp, ok := responseData.(*Config)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetCountriesListParams struct {
	LangCode string // Code for the language used on the client, ISO 639-1 standard
	Hash     int32
}

func (*HelpGetCountriesListParams) CRC() uint32 {
	return 0x735787a8
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetCountriesList(langCode string, hash int32) (HelpCountriesList, error) {
	responseData, err := c.MakeRequest(&HelpGetCountriesListParams{
		Hash:     hash,
		LangCode: langCode,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetCountriesList")
	}

	resp, ok := responseData.(HelpCountriesList)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetDeepLinkInfoParams struct {
	Path string
}

func (*HelpGetDeepLinkInfoParams) CRC() uint32 {
	return 0x3fedc75f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetDeepLinkInfo(path string) (HelpDeepLinkInfo, error) {
	responseData, err := c.MakeRequest(&HelpGetDeepLinkInfoParams{Path: path})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetDeepLinkInfo")
	}

	resp, ok := responseData.(HelpDeepLinkInfo)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetInviteTextParams struct{}

func (*HelpGetInviteTextParams) CRC() uint32 {
	return 0x4d392343
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetInviteText() (*HelpInviteText, error) {
	responseData, err := c.MakeRequest(&HelpGetInviteTextParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetInviteText")
	}

	resp, ok := responseData.(*HelpInviteText)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetNearestDcParams struct{}

func (*HelpGetNearestDcParams) CRC() uint32 {
	return 0x1fb33026
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetNearestDc() (*NearestDc, error) {
	responseData, err := c.MakeRequest(&HelpGetNearestDcParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetNearestDc")
	}

	resp, ok := responseData.(*NearestDc)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetPassportConfigParams struct {
	Hash int32
}

func (*HelpGetPassportConfigParams) CRC() uint32 {
	return 0xc661ad08
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetPassportConfig(hash int32) (HelpPassportConfig, error) {
	responseData, err := c.MakeRequest(&HelpGetPassportConfigParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetPassportConfig")
	}

	resp, ok := responseData.(HelpPassportConfig)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetPromoDataParams struct{}

func (*HelpGetPromoDataParams) CRC() uint32 {
	return 0xc0977421
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetPromoData() (HelpPromoData, error) {
	responseData, err := c.MakeRequest(&HelpGetPromoDataParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetPromoData")
	}

	resp, ok := responseData.(HelpPromoData)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetRecentMeUrlsParams struct {
	Referer string
}

func (*HelpGetRecentMeUrlsParams) CRC() uint32 {
	return 0x3dc0f114
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetRecentMeUrls(referer string) (*HelpRecentMeUrls, error) {
	responseData, err := c.MakeRequest(&HelpGetRecentMeUrlsParams{Referer: referer})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetRecentMeUrls")
	}

	resp, ok := responseData.(*HelpRecentMeUrls)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetSupportParams struct{}

func (*HelpGetSupportParams) CRC() uint32 {
	return 0x9cdf08cd
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetSupport() (*HelpSupport, error) {
	responseData, err := c.MakeRequest(&HelpGetSupportParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetSupport")
	}

	resp, ok := responseData.(*HelpSupport)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetSupportNameParams struct{}

func (*HelpGetSupportNameParams) CRC() uint32 {
	return 0xd360e72c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetSupportName() (*HelpSupportName, error) {
	responseData, err := c.MakeRequest(&HelpGetSupportNameParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetSupportName")
	}

	resp, ok := responseData.(*HelpSupportName)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetTermsOfServiceUpdateParams struct{}

func (*HelpGetTermsOfServiceUpdateParams) CRC() uint32 {
	return 0x2ca51fd1
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetTermsOfServiceUpdate() (HelpTermsOfServiceUpdate, error) {
	responseData, err := c.MakeRequest(&HelpGetTermsOfServiceUpdateParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetTermsOfServiceUpdate")
	}

	resp, ok := responseData.(HelpTermsOfServiceUpdate)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpGetUserInfoParams struct {
	UserID InputUser
}

func (*HelpGetUserInfoParams) CRC() uint32 {
	return 0x38a08d3
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpGetUserInfo(userID InputUser) (HelpUserInfo, error) {
	responseData, err := c.MakeRequest(&HelpGetUserInfoParams{UserID: userID})
	if err != nil {
		return nil, errors.Wrap(err, "sending HelpGetUserInfo")
	}

	resp, ok := responseData.(HelpUserInfo)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpHidePromoDataParams struct {
	Peer InputPeer
}

func (*HelpHidePromoDataParams) CRC() uint32 {
	return 0x1e251c95
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpHidePromoData(peer InputPeer) (bool, error) {
	responseData, err := c.MakeRequest(&HelpHidePromoDataParams{Peer: peer})
	if err != nil {
		return false, errors.Wrap(err, "sending HelpHidePromoData")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpSaveAppLogParams struct {
	Events []*InputAppEvent
}

func (*HelpSaveAppLogParams) CRC() uint32 {
	return 0x6f02f748
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpSaveAppLog(events []*InputAppEvent) (bool, error) {
	responseData, err := c.MakeRequest(&HelpSaveAppLogParams{Events: events})
	if err != nil {
		return false, errors.Wrap(err, "sending HelpSaveAppLog")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type HelpSetBotUpdatesStatusParams struct {
	PendingUpdatesCount int32
	Message             string
}

func (*HelpSetBotUpdatesStatusParams) CRC() uint32 {
	return 0xec22cfcd
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) HelpSetBotUpdatesStatus(pendingUpdatesCount int32, message string) (bool, error) {
	responseData, err := c.MakeRequest(&HelpSetBotUpdatesStatusParams{
		Message:             message,
		PendingUpdatesCount: pendingUpdatesCount,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending HelpSetBotUpdatesStatus")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type LangpackGetDifferenceParams struct {
	LangPack    string // Language pack to use
	LangCode    string // Code for the language used on the client, ISO 639-1 standard
	FromVersion int32
}

func (*LangpackGetDifferenceParams) CRC() uint32 {
	return 0xcd984aa5
}

// Get the participants of a channel
func (c *Client) LangpackGetDifference(langPack, langCode string, fromVersion int32) (*LangPackDifference, error) {
	responseData, err := c.MakeRequest(&LangpackGetDifferenceParams{
		FromVersion: fromVersion,
		LangCode:    langCode,
		LangPack:    langPack,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending LangpackGetDifference")
	}

	resp, ok := responseData.(*LangPackDifference)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type LangpackGetLangPackParams struct {
	LangPack string // Language pack to use
	LangCode string // Code for the language used on the client, ISO 639-1 standard
}

func (*LangpackGetLangPackParams) CRC() uint32 {
	return 0xf2f2330a
}

// Get the participants of a channel
func (c *Client) LangpackGetLangPack(langPack, langCode string) (*LangPackDifference, error) {
	responseData, err := c.MakeRequest(&LangpackGetLangPackParams{
		LangCode: langCode,
		LangPack: langPack,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending LangpackGetLangPack")
	}

	resp, ok := responseData.(*LangPackDifference)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type LangpackGetLanguageParams struct {
	LangPack string // Language pack to use
	LangCode string // Code for the language used on the client, ISO 639-1 standard
}

func (*LangpackGetLanguageParams) CRC() uint32 {
	return 0x6a596502
}

// Get the participants of a channel
func (c *Client) LangpackGetLanguage(langPack, langCode string) (*LangPackLanguage, error) {
	responseData, err := c.MakeRequest(&LangpackGetLanguageParams{
		LangCode: langCode,
		LangPack: langPack,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending LangpackGetLanguage")
	}

	resp, ok := responseData.(*LangPackLanguage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type LangpackGetLanguagesParams struct {
	LangPack string // Language pack to use
}

func (*LangpackGetLanguagesParams) CRC() uint32 {
	return 0x42c6978f
}

// Get the participants of a channel
func (c *Client) LangpackGetLanguages(langPack string) ([]*LangPackLanguage, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&LangpackGetLanguagesParams{LangPack: langPack}, reflect.TypeOf([]*LangPackLanguage{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending LangpackGetLanguages")
	}

	resp, ok := responseData.([]*LangPackLanguage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type LangpackGetStringsParams struct {
	LangPack string // Language pack to use
	LangCode string // Code for the language used on the client, ISO 639-1 standard
	Keys     []string
}

func (*LangpackGetStringsParams) CRC() uint32 {
	return 0xefea3803
}

// Get the participants of a channel
func (c *Client) LangpackGetStrings(langPack, langCode string, keys []string) ([]LangPackString, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&LangpackGetStringsParams{
		Keys:     keys,
		LangCode: langCode,
		LangPack: langPack,
	}, reflect.TypeOf([]LangPackString{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending LangpackGetStrings")
	}

	resp, ok := responseData.([]LangPackString)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesAcceptEncryptionParams struct {
	Peer           *InputEncryptedChat
	GB             []byte
	KeyFingerprint int64
}

func (*MessagesAcceptEncryptionParams) CRC() uint32 {
	return 0x3dbc0415
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesAcceptEncryption(peer *InputEncryptedChat, gB []byte, keyFingerprint int64) (EncryptedChat, error) {
	responseData, err := c.MakeRequest(&MessagesAcceptEncryptionParams{
		GB:             gB,
		KeyFingerprint: keyFingerprint,
		Peer:           peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesAcceptEncryption")
	}

	resp, ok := responseData.(EncryptedChat)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesAcceptURLAuthParams struct {
	WriteAllowed bool `tl:"flag:0,encoded_in_bitflags"`
	Peer         InputPeer
	MsgID        int32 // Message identifier on which a current query depends
	ButtonID     int32
}

func (*MessagesAcceptURLAuthParams) CRC() uint32 {
	return 0xf729ea98
}

func (*MessagesAcceptURLAuthParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesAcceptURLAuth(writeAllowed bool, peer InputPeer, msgID, buttonID int32) (URLAuthResult, error) {
	responseData, err := c.MakeRequest(&MessagesAcceptURLAuthParams{
		ButtonID:     buttonID,
		MsgID:        msgID,
		Peer:         peer,
		WriteAllowed: writeAllowed,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesAcceptURLAuth")
	}

	resp, ok := responseData.(URLAuthResult)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesAddChatUserParams struct {
	ChatID   int32
	UserID   InputUser
	FwdLimit int32
}

func (*MessagesAddChatUserParams) CRC() uint32 {
	return 0xf9a0aa09
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesAddChatUser(chatID int32, userID InputUser, fwdLimit int32) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesAddChatUserParams{
		ChatID:   chatID,
		FwdLimit: fwdLimit,
		UserID:   userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesAddChatUser")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesCheckChatInviteParams struct {
	Hash string
}

func (*MessagesCheckChatInviteParams) CRC() uint32 {
	return 0x3eadb1bb
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesCheckChatInvite(hash string) (ChatInvite, error) {
	responseData, err := c.MakeRequest(&MessagesCheckChatInviteParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesCheckChatInvite")
	}

	resp, ok := responseData.(ChatInvite)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesClearAllDraftsParams struct{}

func (*MessagesClearAllDraftsParams) CRC() uint32 {
	return 0x7e58ee9c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesClearAllDrafts() (bool, error) {
	responseData, err := c.MakeRequest(&MessagesClearAllDraftsParams{})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesClearAllDrafts")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesClearRecentStickersParams struct {
	Attached bool `tl:"flag:0,encoded_in_bitflags"`
}

func (*MessagesClearRecentStickersParams) CRC() uint32 {
	return 0x8999602d
}

func (*MessagesClearRecentStickersParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesClearRecentStickers(attached bool) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesClearRecentStickersParams{Attached: attached})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesClearRecentStickers")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesCreateChatParams struct {
	Users []InputUser
	Title string
}

func (*MessagesCreateChatParams) CRC() uint32 {
	return 0x9cb126e
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesCreateChat(users []InputUser, title string) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesCreateChatParams{
		Title: title,
		Users: users,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesCreateChat")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesDeleteChatUserParams struct {
	ChatID int32
	UserID InputUser
}

func (*MessagesDeleteChatUserParams) CRC() uint32 {
	return 0xe0611f16
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesDeleteChatUser(chatID int32, userID InputUser) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesDeleteChatUserParams{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesDeleteChatUser")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesDeleteHistoryParams struct {
	JustClear bool `tl:"flag:0,encoded_in_bitflags"`
	Revoke    bool `tl:"flag:1,encoded_in_bitflags"`
	Peer      InputPeer
	MaxID     int32
}

func (*MessagesDeleteHistoryParams) CRC() uint32 {
	return 0x1c015b09
}

func (*MessagesDeleteHistoryParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesDeleteHistory(justClear, revoke bool, peer InputPeer, maxID int32) (*MessagesAffectedHistory, error) {
	responseData, err := c.MakeRequest(&MessagesDeleteHistoryParams{
		JustClear: justClear,
		MaxID:     maxID,
		Peer:      peer,
		Revoke:    revoke,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesDeleteHistory")
	}

	resp, ok := responseData.(*MessagesAffectedHistory)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesDeleteMessagesParams struct {
	Revoke bool `tl:"flag:0,encoded_in_bitflags"`
	ID     []int32
}

func (*MessagesDeleteMessagesParams) CRC() uint32 {
	return 0xe58e95d2
}

func (*MessagesDeleteMessagesParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesDeleteMessages(revoke bool, id []int32) (*MessagesAffectedMessages, error) {
	responseData, err := c.MakeRequest(&MessagesDeleteMessagesParams{
		ID:     id,
		Revoke: revoke,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesDeleteMessages")
	}

	resp, ok := responseData.(*MessagesAffectedMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesDeleteScheduledMessagesParams struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesDeleteScheduledMessagesParams) CRC() uint32 {
	return 0x59ae2b16
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesDeleteScheduledMessages(peer InputPeer, id []int32) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesDeleteScheduledMessagesParams{
		ID:   id,
		Peer: peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesDeleteScheduledMessages")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesDiscardEncryptionParams struct {
	ChatID int32
}

func (*MessagesDiscardEncryptionParams) CRC() uint32 {
	return 0xedd923c5
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesDiscardEncryption(chatID int32) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesDiscardEncryptionParams{ChatID: chatID})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesDiscardEncryption")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesEditChatAboutParams struct {
	Peer  InputPeer
	About string
}

func (*MessagesEditChatAboutParams) CRC() uint32 {
	return 0xdef60797
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesEditChatAbout(peer InputPeer, about string) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesEditChatAboutParams{
		About: about,
		Peer:  peer,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesEditChatAbout")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesEditChatAdminParams struct {
	ChatID  int32
	UserID  InputUser
	IsAdmin bool
}

func (*MessagesEditChatAdminParams) CRC() uint32 {
	return 0xa9e69f2e
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesEditChatAdmin(chatID int32, userID InputUser, isAdmin bool) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesEditChatAdminParams{
		ChatID:  chatID,
		IsAdmin: isAdmin,
		UserID:  userID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesEditChatAdmin")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesEditChatDefaultBannedRightsParams struct {
	Peer         InputPeer
	BannedRights *ChatBannedRights
}

func (*MessagesEditChatDefaultBannedRightsParams) CRC() uint32 {
	return 0xa5866b41
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesEditChatDefaultBannedRights(peer InputPeer, bannedRights *ChatBannedRights) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesEditChatDefaultBannedRightsParams{
		BannedRights: bannedRights,
		Peer:         peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesEditChatDefaultBannedRights")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesEditChatPhotoParams struct {
	ChatID int32
	Photo  InputChatPhoto
}

func (*MessagesEditChatPhotoParams) CRC() uint32 {
	return 0xca4c79d8
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesEditChatPhoto(chatID int32, photo InputChatPhoto) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesEditChatPhotoParams{
		ChatID: chatID,
		Photo:  photo,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesEditChatPhoto")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesEditChatTitleParams struct {
	ChatID int32
	Title  string
}

func (*MessagesEditChatTitleParams) CRC() uint32 {
	return 0xdc452855
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesEditChatTitle(chatID int32, title string) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesEditChatTitleParams{
		ChatID: chatID,
		Title:  title,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesEditChatTitle")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesEditInlineBotMessageParams struct {
	NoWebpage   bool `tl:"flag:1,encoded_in_bitflags"`
	ID          *InputBotInlineMessageID
	Message     string          `tl:"flag:11"`
	Media       InputMedia      `tl:"flag:14"`
	ReplyMarkup ReplyMarkup     `tl:"flag:2"`
	Entities    []MessageEntity `tl:"flag:3"`
}

func (*MessagesEditInlineBotMessageParams) CRC() uint32 {
	return 0x83557dba
}

func (*MessagesEditInlineBotMessageParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesEditInlineBotMessage(params *MessagesEditInlineBotMessageParams) (bool, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesEditInlineBotMessage")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesEditMessageParams struct {
	NoWebpage    bool `tl:"flag:1,encoded_in_bitflags"`
	Peer         InputPeer
	ID           int32
	Message      string          `tl:"flag:11"`
	Media        InputMedia      `tl:"flag:14"`
	ReplyMarkup  ReplyMarkup     `tl:"flag:2"`
	Entities     []MessageEntity `tl:"flag:3"`
	ScheduleDate int32           `tl:"flag:15"`
}

func (*MessagesEditMessageParams) CRC() uint32 {
	return 0x48f71778
}

func (*MessagesEditMessageParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesEditMessage(params *MessagesEditMessageParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesEditMessage")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesExportChatInviteParams struct {
	Peer InputPeer
}

func (*MessagesExportChatInviteParams) CRC() uint32 {
	return 0xdf7534c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesExportChatInvite(peer InputPeer) (ExportedChatInvite, error) {
	responseData, err := c.MakeRequest(&MessagesExportChatInviteParams{Peer: peer})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesExportChatInvite")
	}

	resp, ok := responseData.(ExportedChatInvite)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesFaveStickerParams struct {
	ID     InputDocument
	Unfave bool
}

func (*MessagesFaveStickerParams) CRC() uint32 {
	return 0xb9ffc55b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesFaveSticker(id InputDocument, unfave bool) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesFaveStickerParams{
		ID:     id,
		Unfave: unfave,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesFaveSticker")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesForwardMessagesParams struct {
	Silent       bool `tl:"flag:5,encoded_in_bitflags"`
	Background   bool `tl:"flag:6,encoded_in_bitflags"`
	WithMyScore  bool `tl:"flag:8,encoded_in_bitflags"`
	FromPeer     InputPeer
	ID           []int32
	RandomID     []int64
	ToPeer       InputPeer
	ScheduleDate int32 `tl:"flag:10"`
}

func (*MessagesForwardMessagesParams) CRC() uint32 {
	return 0xd9fee60e
}

func (*MessagesForwardMessagesParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesForwardMessages(params *MessagesForwardMessagesParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesForwardMessages")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetAllChatsParams struct {
	ExceptIds []int32
}

func (*MessagesGetAllChatsParams) CRC() uint32 {
	return 0xeba80ff0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetAllChats(exceptIds []int32) (MessagesChats, error) {
	responseData, err := c.MakeRequest(&MessagesGetAllChatsParams{ExceptIds: exceptIds})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetAllChats")
	}

	resp, ok := responseData.(MessagesChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetAllDraftsParams struct{}

func (*MessagesGetAllDraftsParams) CRC() uint32 {
	return 0x6a3f8d65
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetAllDrafts() (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesGetAllDraftsParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetAllDrafts")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetAllStickersParams struct {
	Hash int32
}

func (*MessagesGetAllStickersParams) CRC() uint32 {
	return 0x1c9618b1
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetAllStickers(hash int32) (MessagesAllStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetAllStickersParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetAllStickers")
	}

	resp, ok := responseData.(MessagesAllStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetArchivedStickersParams struct {
	Masks    bool `tl:"flag:0,encoded_in_bitflags"`
	OffsetID int64
	Limit    int32
}

func (*MessagesGetArchivedStickersParams) CRC() uint32 {
	return 0x57f17692
}

func (*MessagesGetArchivedStickersParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetArchivedStickers(masks bool, offsetID int64, limit int32) (*MessagesArchivedStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetArchivedStickersParams{
		Limit:    limit,
		Masks:    masks,
		OffsetID: offsetID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetArchivedStickers")
	}

	resp, ok := responseData.(*MessagesArchivedStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetAttachedStickersParams struct {
	Media InputStickeredMedia
}

func (*MessagesGetAttachedStickersParams) CRC() uint32 {
	return 0xcc5b67cc
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetAttachedStickers(media InputStickeredMedia) ([]StickerSetCovered, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesGetAttachedStickersParams{Media: media}, reflect.TypeOf([]StickerSetCovered{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetAttachedStickers")
	}

	resp, ok := responseData.([]StickerSetCovered)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetBotCallbackAnswerParams struct {
	Game     bool `tl:"flag:1,encoded_in_bitflags"`
	Peer     InputPeer
	MsgID    int32                 // Message identifier on which a current query depends
	Data     []byte                `tl:"flag:0"`
	Password InputCheckPasswordSRP `tl:"flag:2"`
}

func (*MessagesGetBotCallbackAnswerParams) CRC() uint32 {
	return 0x9342ca07
}

func (*MessagesGetBotCallbackAnswerParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetBotCallbackAnswer(params *MessagesGetBotCallbackAnswerParams) (*MessagesBotCallbackAnswer, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetBotCallbackAnswer")
	}

	resp, ok := responseData.(*MessagesBotCallbackAnswer)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetChatsParams struct {
	ID []int32
}

func (*MessagesGetChatsParams) CRC() uint32 {
	return 0x3c6aa187
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetChats(id []int32) (MessagesChats, error) {
	responseData, err := c.MakeRequest(&MessagesGetChatsParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetChats")
	}

	resp, ok := responseData.(MessagesChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetCommonChatsParams struct {
	UserID InputUser
	MaxID  int32
	Limit  int32
}

func (*MessagesGetCommonChatsParams) CRC() uint32 {
	return 0xd0a48c4
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetCommonChats(userID InputUser, maxID, limit int32) (MessagesChats, error) {
	responseData, err := c.MakeRequest(&MessagesGetCommonChatsParams{
		Limit:  limit,
		MaxID:  maxID,
		UserID: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetCommonChats")
	}

	resp, ok := responseData.(MessagesChats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetDhConfigParams struct {
	Version      int32
	RandomLength int32
}

func (*MessagesGetDhConfigParams) CRC() uint32 {
	return 0x26cf8950
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetDhConfig(version, randomLength int32) (MessagesDhConfig, error) {
	responseData, err := c.MakeRequest(&MessagesGetDhConfigParams{
		RandomLength: randomLength,
		Version:      version,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetDhConfig")
	}

	resp, ok := responseData.(MessagesDhConfig)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetDialogFiltersParams struct{}

func (*MessagesGetDialogFiltersParams) CRC() uint32 {
	return 0xf19ed96d
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetDialogFilters() ([]*DialogFilter, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesGetDialogFiltersParams{}, reflect.TypeOf([]*DialogFilter{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetDialogFilters")
	}

	resp, ok := responseData.([]*DialogFilter)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetDialogUnreadMarksParams struct{}

func (*MessagesGetDialogUnreadMarksParams) CRC() uint32 {
	return 0x22e24e22
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetDialogUnreadMarks() ([]DialogPeer, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesGetDialogUnreadMarksParams{}, reflect.TypeOf([]DialogPeer{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetDialogUnreadMarks")
	}

	resp, ok := responseData.([]DialogPeer)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetDialogsParams struct {
	ExcludePinned bool  `tl:"flag:0,encoded_in_bitflags"`
	FolderID      int32 `tl:"flag:1"`
	OffsetDate    int32
	OffsetID      int32
	OffsetPeer    InputPeer
	Limit         int32
	Hash          int32
}

func (*MessagesGetDialogsParams) CRC() uint32 {
	return 0xa0ee3b73
}

func (*MessagesGetDialogsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetDialogs(params *MessagesGetDialogsParams) (MessagesDialogs, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetDialogs")
	}

	resp, ok := responseData.(MessagesDialogs)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetDiscussionMessageParams struct {
	Peer  InputPeer
	MsgID int32 // Message identifier on which a current query depends
}

func (*MessagesGetDiscussionMessageParams) CRC() uint32 {
	return 0x446972fd
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetDiscussionMessage(peer InputPeer, msgID int32) (*MessagesDiscussionMessage, error) {
	responseData, err := c.MakeRequest(&MessagesGetDiscussionMessageParams{
		MsgID: msgID,
		Peer:  peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetDiscussionMessage")
	}

	resp, ok := responseData.(*MessagesDiscussionMessage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetDocumentByHashParams struct {
	SHA256   []byte
	Size     int32
	MimeType string
}

func (*MessagesGetDocumentByHashParams) CRC() uint32 {
	return 0x338e2464
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetDocumentByHash(sha256 []byte, size int32, mimeType string) (Document, error) {
	responseData, err := c.MakeRequest(&MessagesGetDocumentByHashParams{
		MimeType: mimeType,
		SHA256:   sha256,
		Size:     size,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetDocumentByHash")
	}

	resp, ok := responseData.(Document)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetEmojiKeywordsParams struct {
	LangCode string // Code for the language used on the client, ISO 639-1 standard
}

func (*MessagesGetEmojiKeywordsParams) CRC() uint32 {
	return 0x35a0e062
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetEmojiKeywords(langCode string) (*EmojiKeywordsDifference, error) {
	responseData, err := c.MakeRequest(&MessagesGetEmojiKeywordsParams{LangCode: langCode})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetEmojiKeywords")
	}

	resp, ok := responseData.(*EmojiKeywordsDifference)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetEmojiKeywordsDifferenceParams struct {
	LangCode    string // Code for the language used on the client, ISO 639-1 standard
	FromVersion int32
}

func (*MessagesGetEmojiKeywordsDifferenceParams) CRC() uint32 {
	return 0x1508b6af
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetEmojiKeywordsDifference(langCode string, fromVersion int32) (*EmojiKeywordsDifference, error) {
	responseData, err := c.MakeRequest(&MessagesGetEmojiKeywordsDifferenceParams{
		FromVersion: fromVersion,
		LangCode:    langCode,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetEmojiKeywordsDifference")
	}

	resp, ok := responseData.(*EmojiKeywordsDifference)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetEmojiKeywordsLanguagesParams struct {
	LangCodes []string
}

func (*MessagesGetEmojiKeywordsLanguagesParams) CRC() uint32 {
	return 0x4e9963b2
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetEmojiKeywordsLanguages(langCodes []string) ([]*EmojiLanguage, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesGetEmojiKeywordsLanguagesParams{LangCodes: langCodes}, reflect.TypeOf([]*EmojiLanguage{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetEmojiKeywordsLanguages")
	}

	resp, ok := responseData.([]*EmojiLanguage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetEmojiURLParams struct {
	LangCode string // Code for the language used on the client, ISO 639-1 standard
}

func (*MessagesGetEmojiURLParams) CRC() uint32 {
	return 0xd5b10c26
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetEmojiURL(langCode string) (*EmojiURL, error) {
	responseData, err := c.MakeRequest(&MessagesGetEmojiURLParams{LangCode: langCode})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetEmojiURL")
	}

	resp, ok := responseData.(*EmojiURL)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetFavedStickersParams struct {
	Hash int32
}

func (*MessagesGetFavedStickersParams) CRC() uint32 {
	return 0x21ce0b0e
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetFavedStickers(hash int32) (MessagesFavedStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetFavedStickersParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetFavedStickers")
	}

	resp, ok := responseData.(MessagesFavedStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetFeaturedStickersParams struct {
	Hash int32
}

func (*MessagesGetFeaturedStickersParams) CRC() uint32 {
	return 0x2dacca4f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetFeaturedStickers(hash int32) (MessagesFeaturedStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetFeaturedStickersParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetFeaturedStickers")
	}

	resp, ok := responseData.(MessagesFeaturedStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetFullChatParams struct {
	ChatID int32
}

func (*MessagesGetFullChatParams) CRC() uint32 {
	return 0x3b831c66
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetFullChat(chatID int32) (*MessagesChatFull, error) {
	responseData, err := c.MakeRequest(&MessagesGetFullChatParams{ChatID: chatID})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetFullChat")
	}

	resp, ok := responseData.(*MessagesChatFull)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetGameHighScoresParams struct {
	Peer   InputPeer
	ID     int32
	UserID InputUser
}

func (*MessagesGetGameHighScoresParams) CRC() uint32 {
	return 0xe822649d
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetGameHighScores(peer InputPeer, id int32, userID InputUser) (*MessagesHighScores, error) {
	responseData, err := c.MakeRequest(&MessagesGetGameHighScoresParams{
		ID:     id,
		Peer:   peer,
		UserID: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetGameHighScores")
	}

	resp, ok := responseData.(*MessagesHighScores)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetHistoryParams struct {
	Peer       InputPeer
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int32
}

func (*MessagesGetHistoryParams) CRC() uint32 {
	return 0xdcbb8260
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetHistory(params *MessagesGetHistoryParams) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetHistory")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetInlineBotResultsParams struct {
	Bot      InputUser
	Peer     InputPeer
	GeoPoint InputGeoPoint `tl:"flag:0"`
	Query    string        // Query
	Offset   string
}

func (*MessagesGetInlineBotResultsParams) CRC() uint32 {
	return 0x514e999d
}

func (*MessagesGetInlineBotResultsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetInlineBotResults(params *MessagesGetInlineBotResultsParams) (*MessagesBotResults, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetInlineBotResults")
	}

	resp, ok := responseData.(*MessagesBotResults)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetInlineGameHighScoresParams struct {
	ID     *InputBotInlineMessageID
	UserID InputUser
}

func (*MessagesGetInlineGameHighScoresParams) CRC() uint32 {
	return 0xf635e1b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetInlineGameHighScores(id *InputBotInlineMessageID, userID InputUser) (*MessagesHighScores, error) {
	responseData, err := c.MakeRequest(&MessagesGetInlineGameHighScoresParams{
		ID:     id,
		UserID: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetInlineGameHighScores")
	}

	resp, ok := responseData.(*MessagesHighScores)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetMaskStickersParams struct {
	Hash int32
}

func (*MessagesGetMaskStickersParams) CRC() uint32 {
	return 0x65b8c79f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetMaskStickers(hash int32) (MessagesAllStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetMaskStickersParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetMaskStickers")
	}

	resp, ok := responseData.(MessagesAllStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetMessageEditDataParams struct {
	Peer InputPeer
	ID   int32
}

func (*MessagesGetMessageEditDataParams) CRC() uint32 {
	return 0xfda68d36
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetMessageEditData(peer InputPeer, id int32) (*MessagesMessageEditData, error) {
	responseData, err := c.MakeRequest(&MessagesGetMessageEditDataParams{
		ID:   id,
		Peer: peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetMessageEditData")
	}

	resp, ok := responseData.(*MessagesMessageEditData)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetMessagesParams struct {
	ID []InputMessage
}

func (*MessagesGetMessagesParams) CRC() uint32 {
	return 0x63c66506
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetMessages(id []InputMessage) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(&MessagesGetMessagesParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetMessages")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetMessagesViewsParams struct {
	Peer      InputPeer
	ID        []int32
	Increment bool
}

func (*MessagesGetMessagesViewsParams) CRC() uint32 {
	return 0x5784d3e1
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetMessagesViews(peer InputPeer, id []int32, increment bool) (*MessagesMessageViews, error) {
	responseData, err := c.MakeRequest(&MessagesGetMessagesViewsParams{
		ID:        id,
		Increment: increment,
		Peer:      peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetMessagesViews")
	}

	resp, ok := responseData.(*MessagesMessageViews)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetOldFeaturedStickersParams struct {
	Offset int32
	Limit  int32
	Hash   int32
}

func (*MessagesGetOldFeaturedStickersParams) CRC() uint32 {
	return 0x5fe7025b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetOldFeaturedStickers(offset, limit, hash int32) (MessagesFeaturedStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetOldFeaturedStickersParams{
		Hash:   hash,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetOldFeaturedStickers")
	}

	resp, ok := responseData.(MessagesFeaturedStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetOnlinesParams struct {
	Peer InputPeer
}

func (*MessagesGetOnlinesParams) CRC() uint32 {
	return 0x6e2be050
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetOnlines(peer InputPeer) (*ChatOnlines, error) {
	responseData, err := c.MakeRequest(&MessagesGetOnlinesParams{Peer: peer})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetOnlines")
	}

	resp, ok := responseData.(*ChatOnlines)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetPeerDialogsParams struct {
	Peers []InputDialogPeer
}

func (*MessagesGetPeerDialogsParams) CRC() uint32 {
	return 0xe470bcfd
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetPeerDialogs(peers []InputDialogPeer) (*MessagesPeerDialogs, error) {
	responseData, err := c.MakeRequest(&MessagesGetPeerDialogsParams{Peers: peers})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetPeerDialogs")
	}

	resp, ok := responseData.(*MessagesPeerDialogs)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetPeerSettingsParams struct {
	Peer InputPeer
}

func (*MessagesGetPeerSettingsParams) CRC() uint32 {
	return 0x3672e09c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetPeerSettings(peer InputPeer) (*PeerSettings, error) {
	responseData, err := c.MakeRequest(&MessagesGetPeerSettingsParams{Peer: peer})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetPeerSettings")
	}

	resp, ok := responseData.(*PeerSettings)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetPinnedDialogsParams struct {
	FolderID int32
}

func (*MessagesGetPinnedDialogsParams) CRC() uint32 {
	return 0xd6b94df2
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetPinnedDialogs(folderID int32) (*MessagesPeerDialogs, error) {
	responseData, err := c.MakeRequest(&MessagesGetPinnedDialogsParams{FolderID: folderID})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetPinnedDialogs")
	}

	resp, ok := responseData.(*MessagesPeerDialogs)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetPollResultsParams struct {
	Peer  InputPeer
	MsgID int32 // Message identifier on which a current query depends
}

func (*MessagesGetPollResultsParams) CRC() uint32 {
	return 0x73bb643b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetPollResults(peer InputPeer, msgID int32) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesGetPollResultsParams{
		MsgID: msgID,
		Peer:  peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetPollResults")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetPollVotesParams struct {
	Peer   InputPeer
	ID     int32
	Option []byte `tl:"flag:0"`
	Offset string `tl:"flag:1"`
	Limit  int32
}

func (*MessagesGetPollVotesParams) CRC() uint32 {
	return 0xb86e380e
}

func (*MessagesGetPollVotesParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetPollVotes(params *MessagesGetPollVotesParams) (*MessagesVotesList, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetPollVotes")
	}

	resp, ok := responseData.(*MessagesVotesList)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetRecentLocationsParams struct {
	Peer  InputPeer
	Limit int32
	Hash  int32
}

func (*MessagesGetRecentLocationsParams) CRC() uint32 {
	return 0xbbc45b09
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetRecentLocations(peer InputPeer, limit, hash int32) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(&MessagesGetRecentLocationsParams{
		Hash:  hash,
		Limit: limit,
		Peer:  peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetRecentLocations")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetRecentStickersParams struct {
	Attached bool `tl:"flag:0,encoded_in_bitflags"`
	Hash     int32
}

func (*MessagesGetRecentStickersParams) CRC() uint32 {
	return 0x5ea192c9
}

func (*MessagesGetRecentStickersParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetRecentStickers(attached bool, hash int32) (MessagesRecentStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetRecentStickersParams{
		Attached: attached,
		Hash:     hash,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetRecentStickers")
	}

	resp, ok := responseData.(MessagesRecentStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetRepliesParams struct {
	Peer       InputPeer
	MsgID      int32 // Message identifier on which a current query depends
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int32
}

func (*MessagesGetRepliesParams) CRC() uint32 {
	return 0x24b581ba
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetReplies(params *MessagesGetRepliesParams) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetReplies")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetSavedGifsParams struct {
	Hash int32
}

func (*MessagesGetSavedGifsParams) CRC() uint32 {
	return 0x83bf3d52
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetSavedGifs(hash int32) (MessagesSavedGifs, error) {
	responseData, err := c.MakeRequest(&MessagesGetSavedGifsParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetSavedGifs")
	}

	resp, ok := responseData.(MessagesSavedGifs)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetScheduledHistoryParams struct {
	Peer InputPeer
	Hash int32
}

func (*MessagesGetScheduledHistoryParams) CRC() uint32 {
	return 0xe2c2685b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetScheduledHistory(peer InputPeer, hash int32) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(&MessagesGetScheduledHistoryParams{
		Hash: hash,
		Peer: peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetScheduledHistory")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetScheduledMessagesParams struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesGetScheduledMessagesParams) CRC() uint32 {
	return 0xbdbb0464
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetScheduledMessages(peer InputPeer, id []int32) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(&MessagesGetScheduledMessagesParams{
		ID:   id,
		Peer: peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetScheduledMessages")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetSearchCountersParams struct {
	Peer    InputPeer
	Filters []MessagesFilter
}

func (*MessagesGetSearchCountersParams) CRC() uint32 {
	return 0x732eef00
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetSearchCounters(peer InputPeer, filters []MessagesFilter) ([]*MessagesSearchCounter, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesGetSearchCountersParams{
		Filters: filters,
		Peer:    peer,
	}, reflect.TypeOf([]*MessagesSearchCounter{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetSearchCounters")
	}

	resp, ok := responseData.([]*MessagesSearchCounter)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetSplitRangesParams struct{}

func (*MessagesGetSplitRangesParams) CRC() uint32 {
	return 0x1cff7e08
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetSplitRanges() ([]*MessageRange, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesGetSplitRangesParams{}, reflect.TypeOf([]*MessageRange{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetSplitRanges")
	}

	resp, ok := responseData.([]*MessageRange)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetStatsURLParams struct {
	Dark   bool `tl:"flag:0,encoded_in_bitflags"`
	Peer   InputPeer
	Params string // Additional initConnection parameters. For now, only the tz_offset field is supported, for specifying timezone offset in seconds.
}

func (*MessagesGetStatsURLParams) CRC() uint32 {
	return 0x812c2ae6
}

func (*MessagesGetStatsURLParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetStatsURL(dark bool, peer InputPeer, params string) (*StatsURL, error) {
	responseData, err := c.MakeRequest(&MessagesGetStatsURLParams{
		Dark:   dark,
		Params: params,
		Peer:   peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetStatsURL")
	}

	resp, ok := responseData.(*StatsURL)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetStickerSetParams struct {
	Stickerset InputStickerSet
}

func (*MessagesGetStickerSetParams) CRC() uint32 {
	return 0x2619a90e
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetStickerSet(stickerset InputStickerSet) (*MessagesStickerSet, error) {
	responseData, err := c.MakeRequest(&MessagesGetStickerSetParams{Stickerset: stickerset})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetStickerSet")
	}

	resp, ok := responseData.(*MessagesStickerSet)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetStickersParams struct {
	Emoticon string
	Hash     int32
}

func (*MessagesGetStickersParams) CRC() uint32 {
	return 0x43d4f2c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetStickers(emoticon string, hash int32) (MessagesStickers, error) {
	responseData, err := c.MakeRequest(&MessagesGetStickersParams{
		Emoticon: emoticon,
		Hash:     hash,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetStickers")
	}

	resp, ok := responseData.(MessagesStickers)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetSuggestedDialogFiltersParams struct{}

func (*MessagesGetSuggestedDialogFiltersParams) CRC() uint32 {
	return 0xa29cd42c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetSuggestedDialogFilters() ([]*DialogFilterSuggested, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesGetSuggestedDialogFiltersParams{}, reflect.TypeOf([]*DialogFilterSuggested{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetSuggestedDialogFilters")
	}

	resp, ok := responseData.([]*DialogFilterSuggested)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetUnreadMentionsParams struct {
	Peer      InputPeer
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
}

func (*MessagesGetUnreadMentionsParams) CRC() uint32 {
	return 0x46578472
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetUnreadMentions(params *MessagesGetUnreadMentionsParams) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetUnreadMentions")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetWebPageParams struct {
	URL  string
	Hash int32
}

func (*MessagesGetWebPageParams) CRC() uint32 {
	return 0x32ca8f91
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetWebPage(url string, hash int32) (WebPage, error) {
	responseData, err := c.MakeRequest(&MessagesGetWebPageParams{
		Hash: hash,
		URL:  url,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetWebPage")
	}

	resp, ok := responseData.(WebPage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesGetWebPagePreviewParams struct {
	Message  string
	Entities []MessageEntity `tl:"flag:3"`
}

func (*MessagesGetWebPagePreviewParams) CRC() uint32 {
	return 0x8b68b0cc
}

func (*MessagesGetWebPagePreviewParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesGetWebPagePreview(message string, entities []MessageEntity) (MessageMedia, error) {
	responseData, err := c.MakeRequest(&MessagesGetWebPagePreviewParams{
		Entities: entities,
		Message:  message,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesGetWebPagePreview")
	}

	resp, ok := responseData.(MessageMedia)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesHidePeerSettingsBarParams struct {
	Peer InputPeer
}

func (*MessagesHidePeerSettingsBarParams) CRC() uint32 {
	return 0x4facb138
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesHidePeerSettingsBar(peer InputPeer) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesHidePeerSettingsBarParams{Peer: peer})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesHidePeerSettingsBar")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesImportChatInviteParams struct {
	Hash string
}

func (*MessagesImportChatInviteParams) CRC() uint32 {
	return 0x6c50051c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesImportChatInvite(hash string) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesImportChatInviteParams{Hash: hash})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesImportChatInvite")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesInstallStickerSetParams struct {
	Stickerset InputStickerSet
	Archived   bool
}

func (*MessagesInstallStickerSetParams) CRC() uint32 {
	return 0xc78fe460
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesInstallStickerSet(stickerset InputStickerSet, archived bool) (MessagesStickerSetInstallResult, error) {
	responseData, err := c.MakeRequest(&MessagesInstallStickerSetParams{
		Archived:   archived,
		Stickerset: stickerset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesInstallStickerSet")
	}

	resp, ok := responseData.(MessagesStickerSetInstallResult)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesMarkDialogUnreadParams struct {
	Unread bool `tl:"flag:0,encoded_in_bitflags"`
	Peer   InputDialogPeer
}

func (*MessagesMarkDialogUnreadParams) CRC() uint32 {
	return 0xc286d98f
}

func (*MessagesMarkDialogUnreadParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesMarkDialogUnread(unread bool, peer InputDialogPeer) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesMarkDialogUnreadParams{
		Peer:   peer,
		Unread: unread,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesMarkDialogUnread")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesMigrateChatParams struct {
	ChatID int32
}

func (*MessagesMigrateChatParams) CRC() uint32 {
	return 0x15a3b8e3
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesMigrateChat(chatID int32) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesMigrateChatParams{ChatID: chatID})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesMigrateChat")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReadDiscussionParams struct {
	Peer      InputPeer
	MsgID     int32 // Message identifier on which a current query depends
	ReadMaxID int32
}

func (*MessagesReadDiscussionParams) CRC() uint32 {
	return 0xf731a9f4
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReadDiscussion(peer InputPeer, msgID, readMaxID int32) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReadDiscussionParams{
		MsgID:     msgID,
		Peer:      peer,
		ReadMaxID: readMaxID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReadDiscussion")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReadEncryptedHistoryParams struct {
	Peer    *InputEncryptedChat
	MaxDate int32
}

func (*MessagesReadEncryptedHistoryParams) CRC() uint32 {
	return 0x7f4b690a
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReadEncryptedHistory(peer *InputEncryptedChat, maxDate int32) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReadEncryptedHistoryParams{
		MaxDate: maxDate,
		Peer:    peer,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReadEncryptedHistory")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReadFeaturedStickersParams struct {
	ID []int64
}

func (*MessagesReadFeaturedStickersParams) CRC() uint32 {
	return 0x5b118126
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReadFeaturedStickers(id []int64) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReadFeaturedStickersParams{ID: id})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReadFeaturedStickers")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReadHistoryParams struct {
	Peer  InputPeer
	MaxID int32
}

func (*MessagesReadHistoryParams) CRC() uint32 {
	return 0xe306d3a
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReadHistory(peer InputPeer, maxID int32) (*MessagesAffectedMessages, error) {
	responseData, err := c.MakeRequest(&MessagesReadHistoryParams{
		MaxID: maxID,
		Peer:  peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesReadHistory")
	}

	resp, ok := responseData.(*MessagesAffectedMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReadMentionsParams struct {
	Peer InputPeer
}

func (*MessagesReadMentionsParams) CRC() uint32 {
	return 0xf0189d3
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReadMentions(peer InputPeer) (*MessagesAffectedHistory, error) {
	responseData, err := c.MakeRequest(&MessagesReadMentionsParams{Peer: peer})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesReadMentions")
	}

	resp, ok := responseData.(*MessagesAffectedHistory)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReadMessageContentsParams struct {
	ID []int32
}

func (*MessagesReadMessageContentsParams) CRC() uint32 {
	return 0x36a73f77
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReadMessageContents(id []int32) (*MessagesAffectedMessages, error) {
	responseData, err := c.MakeRequest(&MessagesReadMessageContentsParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesReadMessageContents")
	}

	resp, ok := responseData.(*MessagesAffectedMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReceivedMessagesParams struct {
	MaxID int32
}

func (*MessagesReceivedMessagesParams) CRC() uint32 {
	return 0x5a954c0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReceivedMessages(maxID int32) ([]*ReceivedNotifyMessage, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesReceivedMessagesParams{MaxID: maxID}, reflect.TypeOf([]*ReceivedNotifyMessage{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesReceivedMessages")
	}

	resp, ok := responseData.([]*ReceivedNotifyMessage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReceivedQueueParams struct {
	MaxQts int32
}

func (*MessagesReceivedQueueParams) CRC() uint32 {
	return 0x55a5bb66
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReceivedQueue(maxQts int32) ([]int64, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&MessagesReceivedQueueParams{MaxQts: maxQts}, reflect.TypeOf([]int64{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesReceivedQueue")
	}

	resp, ok := responseData.([]int64)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReorderPinnedDialogsParams struct {
	Force    bool `tl:"flag:0,encoded_in_bitflags"`
	FolderID int32
	Order    []InputDialogPeer
}

func (*MessagesReorderPinnedDialogsParams) CRC() uint32 {
	return 0x3b1adf37
}

func (*MessagesReorderPinnedDialogsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReorderPinnedDialogs(force bool, folderID int32, order []InputDialogPeer) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReorderPinnedDialogsParams{
		FolderID: folderID,
		Force:    force,
		Order:    order,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReorderPinnedDialogs")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReorderStickerSetsParams struct {
	Masks bool `tl:"flag:0,encoded_in_bitflags"`
	Order []int64
}

func (*MessagesReorderStickerSetsParams) CRC() uint32 {
	return 0x78337739
}

func (*MessagesReorderStickerSetsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReorderStickerSets(masks bool, order []int64) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReorderStickerSetsParams{
		Masks: masks,
		Order: order,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReorderStickerSets")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReportParams struct {
	Peer   InputPeer
	ID     []int32
	Reason ReportReason
}

func (*MessagesReportParams) CRC() uint32 {
	return 0xbd82b658
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReport(peer InputPeer, id []int32, reason ReportReason) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReportParams{
		ID:     id,
		Peer:   peer,
		Reason: reason,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReport")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReportEncryptedSpamParams struct {
	Peer *InputEncryptedChat
}

func (*MessagesReportEncryptedSpamParams) CRC() uint32 {
	return 0x4b0c8c0f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReportEncryptedSpam(peer *InputEncryptedChat) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReportEncryptedSpamParams{Peer: peer})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReportEncryptedSpam")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesReportSpamParams struct {
	Peer InputPeer
}

func (*MessagesReportSpamParams) CRC() uint32 {
	return 0xcf1592db
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesReportSpam(peer InputPeer) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesReportSpamParams{Peer: peer})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesReportSpam")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesRequestEncryptionParams struct {
	UserID   InputUser
	RandomID int32
	GA       []byte
}

func (*MessagesRequestEncryptionParams) CRC() uint32 {
	return 0xf64daf43
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesRequestEncryption(userID InputUser, randomID int32, gA []byte) (EncryptedChat, error) {
	responseData, err := c.MakeRequest(&MessagesRequestEncryptionParams{
		GA:       gA,
		RandomID: randomID,
		UserID:   userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesRequestEncryption")
	}

	resp, ok := responseData.(EncryptedChat)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesRequestURLAuthParams struct {
	Peer     InputPeer
	MsgID    int32 // Message identifier on which a current query depends
	ButtonID int32
}

func (*MessagesRequestURLAuthParams) CRC() uint32 {
	return 0xe33f5613
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesRequestURLAuth(peer InputPeer, msgID, buttonID int32) (URLAuthResult, error) {
	responseData, err := c.MakeRequest(&MessagesRequestURLAuthParams{
		ButtonID: buttonID,
		MsgID:    msgID,
		Peer:     peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesRequestURLAuth")
	}

	resp, ok := responseData.(URLAuthResult)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSaveDraftParams struct {
	NoWebpage    bool  `tl:"flag:1,encoded_in_bitflags"`
	ReplyToMsgID int32 `tl:"flag:0"`
	Peer         InputPeer
	Message      string
	Entities     []MessageEntity `tl:"flag:3"`
}

func (*MessagesSaveDraftParams) CRC() uint32 {
	return 0xbc39e14b
}

func (*MessagesSaveDraftParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSaveDraft(params *MessagesSaveDraftParams) (bool, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSaveDraft")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSaveGifParams struct {
	ID     InputDocument
	Unsave bool
}

func (*MessagesSaveGifParams) CRC() uint32 {
	return 0x327a30cb
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSaveGif(id InputDocument, unsave bool) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesSaveGifParams{
		ID:     id,
		Unsave: unsave,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSaveGif")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSaveRecentStickerParams struct {
	Attached bool `tl:"flag:0,encoded_in_bitflags"`
	ID       InputDocument
	Unsave   bool
}

func (*MessagesSaveRecentStickerParams) CRC() uint32 {
	return 0x392718f8
}

func (*MessagesSaveRecentStickerParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSaveRecentSticker(attached bool, id InputDocument, unsave bool) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesSaveRecentStickerParams{
		Attached: attached,
		ID:       id,
		Unsave:   unsave,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSaveRecentSticker")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSearchParams struct {
	Peer      InputPeer
	Q         string
	FromID    InputPeer `tl:"flag:0"`
	TopMsgID  int32     `tl:"flag:1"`
	Filter    MessagesFilter
	MinDate   int32
	MaxDate   int32
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
	Hash      int32
}

func (*MessagesSearchParams) CRC() uint32 {
	return 0xc352eec
}

func (*MessagesSearchParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSearch(params *MessagesSearchParams) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSearch")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSearchGlobalParams struct {
	FolderID   int32 `tl:"flag:0"`
	Q          string
	Filter     MessagesFilter
	MinDate    int32
	MaxDate    int32
	OffsetRate int32
	OffsetPeer InputPeer
	OffsetID   int32
	Limit      int32
}

func (*MessagesSearchGlobalParams) CRC() uint32 {
	return 0x4bc6589a
}

func (*MessagesSearchGlobalParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSearchGlobal(params *MessagesSearchGlobalParams) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSearchGlobal")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSearchStickerSetsParams struct {
	ExcludeFeatured bool `tl:"flag:0,encoded_in_bitflags"`
	Q               string
	Hash            int32
}

func (*MessagesSearchStickerSetsParams) CRC() uint32 {
	return 0xc2b7d08b
}

func (*MessagesSearchStickerSetsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSearchStickerSets(excludeFeatured bool, q string, hash int32) (MessagesFoundStickerSets, error) {
	responseData, err := c.MakeRequest(&MessagesSearchStickerSetsParams{
		ExcludeFeatured: excludeFeatured,
		Hash:            hash,
		Q:               q,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSearchStickerSets")
	}

	resp, ok := responseData.(MessagesFoundStickerSets)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendEncryptedParams struct {
	Silent   bool `tl:"flag:0,encoded_in_bitflags"`
	Peer     *InputEncryptedChat
	RandomID int64
	Data     []byte
}

func (*MessagesSendEncryptedParams) CRC() uint32 {
	return 0x44fa7a15
}

func (*MessagesSendEncryptedParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendEncrypted(silent bool, peer *InputEncryptedChat, randomID int64, data []byte) (MessagesSentEncryptedMessage, error) {
	responseData, err := c.MakeRequest(&MessagesSendEncryptedParams{
		Data:     data,
		Peer:     peer,
		RandomID: randomID,
		Silent:   silent,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendEncrypted")
	}

	resp, ok := responseData.(MessagesSentEncryptedMessage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendEncryptedFileParams struct {
	Silent   bool `tl:"flag:0,encoded_in_bitflags"`
	Peer     *InputEncryptedChat
	RandomID int64
	Data     []byte
	File     InputEncryptedFile
}

func (*MessagesSendEncryptedFileParams) CRC() uint32 {
	return 0x5559481d
}

func (*MessagesSendEncryptedFileParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendEncryptedFile(params *MessagesSendEncryptedFileParams) (MessagesSentEncryptedMessage, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendEncryptedFile")
	}

	resp, ok := responseData.(MessagesSentEncryptedMessage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendEncryptedServiceParams struct {
	Peer     *InputEncryptedChat
	RandomID int64
	Data     []byte
}

func (*MessagesSendEncryptedServiceParams) CRC() uint32 {
	return 0x32d439a4
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendEncryptedService(peer *InputEncryptedChat, randomID int64, data []byte) (MessagesSentEncryptedMessage, error) {
	responseData, err := c.MakeRequest(&MessagesSendEncryptedServiceParams{
		Data:     data,
		Peer:     peer,
		RandomID: randomID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendEncryptedService")
	}

	resp, ok := responseData.(MessagesSentEncryptedMessage)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendInlineBotResultParams struct {
	Silent       bool `tl:"flag:5,encoded_in_bitflags"`
	Background   bool `tl:"flag:6,encoded_in_bitflags"`
	ClearDraft   bool `tl:"flag:7,encoded_in_bitflags"`
	HideVia      bool `tl:"flag:11,encoded_in_bitflags"`
	Peer         InputPeer
	ReplyToMsgID int32 `tl:"flag:0"`
	RandomID     int64
	QueryID      int64
	ID           string
	ScheduleDate int32 `tl:"flag:10"`
}

func (*MessagesSendInlineBotResultParams) CRC() uint32 {
	return 0x220815b0
}

func (*MessagesSendInlineBotResultParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendInlineBotResult(params *MessagesSendInlineBotResultParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendInlineBotResult")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendMediaParams struct {
	Silent       bool `tl:"flag:5,encoded_in_bitflags"`
	Background   bool `tl:"flag:6,encoded_in_bitflags"`
	ClearDraft   bool `tl:"flag:7,encoded_in_bitflags"`
	Peer         InputPeer
	ReplyToMsgID int32 `tl:"flag:0"`
	Media        InputMedia
	Message      string
	RandomID     int64
	ReplyMarkup  ReplyMarkup     `tl:"flag:2"`
	Entities     []MessageEntity `tl:"flag:3"`
	ScheduleDate int32           `tl:"flag:10"`
}

func (*MessagesSendMediaParams) CRC() uint32 {
	return 0x3491eba9
}

func (*MessagesSendMediaParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendMedia(params *MessagesSendMediaParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendMedia")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendMessageParams struct {
	NoWebpage    bool `tl:"flag:1,encoded_in_bitflags"`
	Silent       bool `tl:"flag:5,encoded_in_bitflags"`
	Background   bool `tl:"flag:6,encoded_in_bitflags"`
	ClearDraft   bool `tl:"flag:7,encoded_in_bitflags"`
	Peer         InputPeer
	ReplyToMsgID int32 `tl:"flag:0"`
	Message      string
	RandomID     int64
	ReplyMarkup  ReplyMarkup     `tl:"flag:2"`
	Entities     []MessageEntity `tl:"flag:3"`
	ScheduleDate int32           `tl:"flag:10"`
}

func (*MessagesSendMessageParams) CRC() uint32 {
	return 0x520c3870
}

func (*MessagesSendMessageParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendMessage(params *MessagesSendMessageParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendMessage")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendMultiMediaParams struct {
	Silent       bool `tl:"flag:5,encoded_in_bitflags"`
	Background   bool `tl:"flag:6,encoded_in_bitflags"`
	ClearDraft   bool `tl:"flag:7,encoded_in_bitflags"`
	Peer         InputPeer
	ReplyToMsgID int32 `tl:"flag:0"`
	MultiMedia   []*InputSingleMedia
	ScheduleDate int32 `tl:"flag:10"`
}

func (*MessagesSendMultiMediaParams) CRC() uint32 {
	return 0xcc0110cb
}

func (*MessagesSendMultiMediaParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendMultiMedia(params *MessagesSendMultiMediaParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendMultiMedia")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendScheduledMessagesParams struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesSendScheduledMessagesParams) CRC() uint32 {
	return 0xbd38850a
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendScheduledMessages(peer InputPeer, id []int32) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesSendScheduledMessagesParams{
		ID:   id,
		Peer: peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendScheduledMessages")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendScreenshotNotificationParams struct {
	Peer         InputPeer
	ReplyToMsgID int32
	RandomID     int64
}

func (*MessagesSendScreenshotNotificationParams) CRC() uint32 {
	return 0xc97df020
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendScreenshotNotification(peer InputPeer, replyToMsgID int32, randomID int64) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesSendScreenshotNotificationParams{
		Peer:         peer,
		RandomID:     randomID,
		ReplyToMsgID: replyToMsgID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendScreenshotNotification")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSendVoteParams struct {
	Peer    InputPeer
	MsgID   int32 // Message identifier on which a current query depends
	Options [][]byte
}

func (*MessagesSendVoteParams) CRC() uint32 {
	return 0x10ea6184
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSendVote(peer InputPeer, msgID int32, options [][]byte) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesSendVoteParams{
		MsgID:   msgID,
		Options: options,
		Peer:    peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSendVote")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetBotCallbackAnswerParams struct {
	Alert     bool `tl:"flag:1,encoded_in_bitflags"`
	QueryID   int64
	Message   string `tl:"flag:0"`
	URL       string `tl:"flag:2"`
	CacheTime int32
}

func (*MessagesSetBotCallbackAnswerParams) CRC() uint32 {
	return 0xd58f130a
}

func (*MessagesSetBotCallbackAnswerParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetBotCallbackAnswer(params *MessagesSetBotCallbackAnswerParams) (bool, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSetBotCallbackAnswer")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetBotPrecheckoutResultsParams struct {
	Success bool `tl:"flag:1,encoded_in_bitflags"`
	QueryID int64
	Error   string `tl:"flag:0"`
}

func (*MessagesSetBotPrecheckoutResultsParams) CRC() uint32 {
	return 0x9c2dd95
}

func (*MessagesSetBotPrecheckoutResultsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetBotPrecheckoutResults(success bool, queryID int64, error string) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesSetBotPrecheckoutResultsParams{
		Error:   error,
		QueryID: queryID,
		Success: success,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSetBotPrecheckoutResults")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetBotShippingResultsParams struct {
	QueryID         int64
	Error           string            `tl:"flag:0"`
	ShippingOptions []*ShippingOption `tl:"flag:1"`
}

func (*MessagesSetBotShippingResultsParams) CRC() uint32 {
	return 0xe5f672fa
}

func (*MessagesSetBotShippingResultsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetBotShippingResults(queryID int64, error string, shippingOptions []*ShippingOption) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesSetBotShippingResultsParams{
		Error:           error,
		QueryID:         queryID,
		ShippingOptions: shippingOptions,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSetBotShippingResults")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetEncryptedTypingParams struct {
	Peer   *InputEncryptedChat
	Typing bool
}

func (*MessagesSetEncryptedTypingParams) CRC() uint32 {
	return 0x791451ed
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetEncryptedTyping(peer *InputEncryptedChat, typing bool) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesSetEncryptedTypingParams{
		Peer:   peer,
		Typing: typing,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSetEncryptedTyping")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetGameScoreParams struct {
	EditMessage bool `tl:"flag:0,encoded_in_bitflags"`
	Force       bool `tl:"flag:1,encoded_in_bitflags"`
	Peer        InputPeer
	ID          int32
	UserID      InputUser
	Score       int32
}

func (*MessagesSetGameScoreParams) CRC() uint32 {
	return 0x8ef8ecc0
}

func (*MessagesSetGameScoreParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetGameScore(params *MessagesSetGameScoreParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesSetGameScore")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetInlineBotResultsParams struct {
	Gallery    bool `tl:"flag:0,encoded_in_bitflags"`
	Private    bool `tl:"flag:1,encoded_in_bitflags"`
	QueryID    int64
	Results    []InputBotInlineResult
	CacheTime  int32
	NextOffset string             `tl:"flag:2"`
	SwitchPm   *InlineBotSwitchPm `tl:"flag:3"`
}

func (*MessagesSetInlineBotResultsParams) CRC() uint32 {
	return 0xeb5ea206
}

func (*MessagesSetInlineBotResultsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetInlineBotResults(params *MessagesSetInlineBotResultsParams) (bool, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSetInlineBotResults")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetInlineGameScoreParams struct {
	EditMessage bool `tl:"flag:0,encoded_in_bitflags"`
	Force       bool `tl:"flag:1,encoded_in_bitflags"`
	ID          *InputBotInlineMessageID
	UserID      InputUser
	Score       int32
}

func (*MessagesSetInlineGameScoreParams) CRC() uint32 {
	return 0x15ad9f64
}

func (*MessagesSetInlineGameScoreParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetInlineGameScore(params *MessagesSetInlineGameScoreParams) (bool, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSetInlineGameScore")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesSetTypingParams struct {
	Peer     InputPeer
	TopMsgID int32 `tl:"flag:0"`
	Action   SendMessageAction
}

func (*MessagesSetTypingParams) CRC() uint32 {
	return 0x58943ee2
}

func (*MessagesSetTypingParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesSetTyping(peer InputPeer, topMsgID int32, action SendMessageAction) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesSetTypingParams{
		Action:   action,
		Peer:     peer,
		TopMsgID: topMsgID,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesSetTyping")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesStartBotParams struct {
	Bot        InputUser
	Peer       InputPeer
	RandomID   int64
	StartParam string
}

func (*MessagesStartBotParams) CRC() uint32 {
	return 0xe6df7378
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesStartBot(bot InputUser, peer InputPeer, randomID int64, startParam string) (Updates, error) {
	responseData, err := c.MakeRequest(&MessagesStartBotParams{
		Bot:        bot,
		Peer:       peer,
		RandomID:   randomID,
		StartParam: startParam,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesStartBot")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesToggleDialogPinParams struct {
	Pinned bool `tl:"flag:0,encoded_in_bitflags"`
	Peer   InputDialogPeer
}

func (*MessagesToggleDialogPinParams) CRC() uint32 {
	return 0xa731e257
}

func (*MessagesToggleDialogPinParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesToggleDialogPin(pinned bool, peer InputDialogPeer) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesToggleDialogPinParams{
		Peer:   peer,
		Pinned: pinned,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesToggleDialogPin")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesToggleStickerSetsParams struct {
	Uninstall   bool `tl:"flag:0,encoded_in_bitflags"`
	Archive     bool `tl:"flag:1,encoded_in_bitflags"`
	Unarchive   bool `tl:"flag:2,encoded_in_bitflags"`
	Stickersets []InputStickerSet
}

func (*MessagesToggleStickerSetsParams) CRC() uint32 {
	return 0xb5052fea
}

func (*MessagesToggleStickerSetsParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesToggleStickerSets(uninstall, archive, unarchive bool, stickersets []InputStickerSet) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesToggleStickerSetsParams{
		Archive:     archive,
		Stickersets: stickersets,
		Unarchive:   unarchive,
		Uninstall:   uninstall,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesToggleStickerSets")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesUninstallStickerSetParams struct {
	Stickerset InputStickerSet
}

func (*MessagesUninstallStickerSetParams) CRC() uint32 {
	return 0xf96e55de
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesUninstallStickerSet(stickerset InputStickerSet) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesUninstallStickerSetParams{Stickerset: stickerset})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesUninstallStickerSet")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesUnpinAllMessagesParams struct {
	Peer InputPeer
}

func (*MessagesUnpinAllMessagesParams) CRC() uint32 {
	return 0xf025bc8b
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesUnpinAllMessages(peer InputPeer) (*MessagesAffectedHistory, error) {
	responseData, err := c.MakeRequest(&MessagesUnpinAllMessagesParams{Peer: peer})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesUnpinAllMessages")
	}

	resp, ok := responseData.(*MessagesAffectedHistory)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesUpdateDialogFilterParams struct {
	ID     int32
	Filter *DialogFilter `tl:"flag:0"`
}

func (*MessagesUpdateDialogFilterParams) CRC() uint32 {
	return 0x1ad4a04a
}

func (*MessagesUpdateDialogFilterParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesUpdateDialogFilter(id int32, filter *DialogFilter) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesUpdateDialogFilterParams{
		Filter: filter,
		ID:     id,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesUpdateDialogFilter")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesUpdateDialogFiltersOrderParams struct {
	Order []int32
}

func (*MessagesUpdateDialogFiltersOrderParams) CRC() uint32 {
	return 0xc563c1e4
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesUpdateDialogFiltersOrder(order []int32) (bool, error) {
	responseData, err := c.MakeRequest(&MessagesUpdateDialogFiltersOrderParams{Order: order})
	if err != nil {
		return false, errors.Wrap(err, "sending MessagesUpdateDialogFiltersOrder")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesUpdatePinnedMessageParams struct {
	Silent    bool `tl:"flag:0,encoded_in_bitflags"`
	Unpin     bool `tl:"flag:1,encoded_in_bitflags"`
	PmOneside bool `tl:"flag:2,encoded_in_bitflags"`
	Peer      InputPeer
	ID        int32
}

func (*MessagesUpdatePinnedMessageParams) CRC() uint32 {
	return 0xd2aaf7ec
}

func (*MessagesUpdatePinnedMessageParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesUpdatePinnedMessage(params *MessagesUpdatePinnedMessageParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesUpdatePinnedMessage")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesUploadEncryptedFileParams struct {
	Peer *InputEncryptedChat
	File InputEncryptedFile
}

func (*MessagesUploadEncryptedFileParams) CRC() uint32 {
	return 0x5057c497
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesUploadEncryptedFile(peer *InputEncryptedChat, file InputEncryptedFile) (EncryptedFile, error) {
	responseData, err := c.MakeRequest(&MessagesUploadEncryptedFileParams{
		File: file,
		Peer: peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesUploadEncryptedFile")
	}

	resp, ok := responseData.(EncryptedFile)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type MessagesUploadMediaParams struct {
	Peer  InputPeer
	Media InputMedia
}

func (*MessagesUploadMediaParams) CRC() uint32 {
	return 0x519bc2b1
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) MessagesUploadMedia(peer InputPeer, media InputMedia) (MessageMedia, error) {
	responseData, err := c.MakeRequest(&MessagesUploadMediaParams{
		Media: media,
		Peer:  peer,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending MessagesUploadMedia")
	}

	resp, ok := responseData.(MessageMedia)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PaymentsClearSavedInfoParams struct {
	Credentials bool `tl:"flag:0,encoded_in_bitflags"` // Encrypted values
	Info        bool `tl:"flag:1,encoded_in_bitflags"`
}

func (*PaymentsClearSavedInfoParams) CRC() uint32 {
	return 0xd83d70c1
}

func (*PaymentsClearSavedInfoParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) PaymentsClearSavedInfo(credentials, info bool) (bool, error) {
	responseData, err := c.MakeRequest(&PaymentsClearSavedInfoParams{
		Credentials: credentials,
		Info:        info,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending PaymentsClearSavedInfo")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PaymentsGetBankCardDataParams struct {
	Number string
}

func (*PaymentsGetBankCardDataParams) CRC() uint32 {
	return 0x2e79d779
}

// Get the participants of a channel
func (c *Client) PaymentsGetBankCardData(number string) (*PaymentsBankCardData, error) {
	responseData, err := c.MakeRequest(&PaymentsGetBankCardDataParams{Number: number})
	if err != nil {
		return nil, errors.Wrap(err, "sending PaymentsGetBankCardData")
	}

	resp, ok := responseData.(*PaymentsBankCardData)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PaymentsGetPaymentFormParams struct {
	MsgID int32 // Message identifier on which a current query depends
}

func (*PaymentsGetPaymentFormParams) CRC() uint32 {
	return 0x99f09745
}

// Get the participants of a channel
func (c *Client) PaymentsGetPaymentForm(msgID int32) (*PaymentsPaymentForm, error) {
	responseData, err := c.MakeRequest(&PaymentsGetPaymentFormParams{MsgID: msgID})
	if err != nil {
		return nil, errors.Wrap(err, "sending PaymentsGetPaymentForm")
	}

	resp, ok := responseData.(*PaymentsPaymentForm)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PaymentsGetPaymentReceiptParams struct {
	MsgID int32 // Message identifier on which a current query depends
}

func (*PaymentsGetPaymentReceiptParams) CRC() uint32 {
	return 0xa092a980
}

// Get the participants of a channel
func (c *Client) PaymentsGetPaymentReceipt(msgID int32) (*PaymentsPaymentReceipt, error) {
	responseData, err := c.MakeRequest(&PaymentsGetPaymentReceiptParams{MsgID: msgID})
	if err != nil {
		return nil, errors.Wrap(err, "sending PaymentsGetPaymentReceipt")
	}

	resp, ok := responseData.(*PaymentsPaymentReceipt)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PaymentsGetSavedInfoParams struct{}

func (*PaymentsGetSavedInfoParams) CRC() uint32 {
	return 0x227d824b
}

// Get the participants of a channel
func (c *Client) PaymentsGetSavedInfo() (*PaymentsSavedInfo, error) {
	responseData, err := c.MakeRequest(&PaymentsGetSavedInfoParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending PaymentsGetSavedInfo")
	}

	resp, ok := responseData.(*PaymentsSavedInfo)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PaymentsSendPaymentFormParams struct {
	MsgID            int32                   // Message identifier on which a current query depends
	RequestedInfoID  string                  `tl:"flag:0"`
	ShippingOptionID string                  `tl:"flag:1"`
	Credentials      InputPaymentCredentials // Encrypted values
}

func (*PaymentsSendPaymentFormParams) CRC() uint32 {
	return 0x2b8879b3
}

func (*PaymentsSendPaymentFormParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) PaymentsSendPaymentForm(msgID int32, requestedInfoID, shippingOptionID string, credentials InputPaymentCredentials) (PaymentsPaymentResult, error) {
	responseData, err := c.MakeRequest(&PaymentsSendPaymentFormParams{
		Credentials:      credentials,
		MsgID:            msgID,
		RequestedInfoID:  requestedInfoID,
		ShippingOptionID: shippingOptionID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending PaymentsSendPaymentForm")
	}

	resp, ok := responseData.(PaymentsPaymentResult)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PaymentsValidateRequestedInfoParams struct {
	Save  bool  `tl:"flag:0,encoded_in_bitflags"`
	MsgID int32 // Message identifier on which a current query depends
	Info  *PaymentRequestedInfo
}

func (*PaymentsValidateRequestedInfoParams) CRC() uint32 {
	return 0x770a8e74
}

func (*PaymentsValidateRequestedInfoParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) PaymentsValidateRequestedInfo(save bool, msgID int32, info *PaymentRequestedInfo) (*PaymentsValidatedRequestedInfo, error) {
	responseData, err := c.MakeRequest(&PaymentsValidateRequestedInfoParams{
		Info:  info,
		MsgID: msgID,
		Save:  save,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending PaymentsValidateRequestedInfo")
	}

	resp, ok := responseData.(*PaymentsValidatedRequestedInfo)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneAcceptCallParams struct {
	Peer     *InputPhoneCall
	GB       []byte
	Protocol *PhoneCallProtocol
}

func (*PhoneAcceptCallParams) CRC() uint32 {
	return 0x3bd2b4a0
}

// Get the participants of a channel
func (c *Client) PhoneAcceptCall(peer *InputPhoneCall, gB []byte, protocol *PhoneCallProtocol) (*PhonePhoneCall, error) {
	responseData, err := c.MakeRequest(&PhoneAcceptCallParams{
		GB:       gB,
		Peer:     peer,
		Protocol: protocol,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending PhoneAcceptCall")
	}

	resp, ok := responseData.(*PhonePhoneCall)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneConfirmCallParams struct {
	Peer           *InputPhoneCall
	GA             []byte
	KeyFingerprint int64
	Protocol       *PhoneCallProtocol
}

func (*PhoneConfirmCallParams) CRC() uint32 {
	return 0x2efe1722
}

// Get the participants of a channel
func (c *Client) PhoneConfirmCall(peer *InputPhoneCall, gA []byte, keyFingerprint int64, protocol *PhoneCallProtocol) (*PhonePhoneCall, error) {
	responseData, err := c.MakeRequest(&PhoneConfirmCallParams{
		GA:             gA,
		KeyFingerprint: keyFingerprint,
		Peer:           peer,
		Protocol:       protocol,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending PhoneConfirmCall")
	}

	resp, ok := responseData.(*PhonePhoneCall)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneDiscardCallParams struct {
	Video        bool `tl:"flag:0,encoded_in_bitflags"`
	Peer         *InputPhoneCall
	Duration     int32
	Reason       PhoneCallDiscardReason
	ConnectionID int64
}

func (*PhoneDiscardCallParams) CRC() uint32 {
	return 0xb2cbc1c0
}

func (*PhoneDiscardCallParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) PhoneDiscardCall(params *PhoneDiscardCallParams) (Updates, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending PhoneDiscardCall")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneGetCallConfigParams struct{}

func (*PhoneGetCallConfigParams) CRC() uint32 {
	return 0x55451fa9
}

// Get the participants of a channel
func (c *Client) PhoneGetCallConfig() (*DataJson, error) {
	responseData, err := c.MakeRequest(&PhoneGetCallConfigParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending PhoneGetCallConfig")
	}

	resp, ok := responseData.(*DataJson)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneReceivedCallParams struct {
	Peer *InputPhoneCall
}

func (*PhoneReceivedCallParams) CRC() uint32 {
	return 0x17d54f61
}

// Get the participants of a channel
func (c *Client) PhoneReceivedCall(peer *InputPhoneCall) (bool, error) {
	responseData, err := c.MakeRequest(&PhoneReceivedCallParams{Peer: peer})
	if err != nil {
		return false, errors.Wrap(err, "sending PhoneReceivedCall")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneRequestCallParams struct {
	Video    bool `tl:"flag:0,encoded_in_bitflags"`
	UserID   InputUser
	RandomID int32
	GAHash   []byte
	Protocol *PhoneCallProtocol
}

func (*PhoneRequestCallParams) CRC() uint32 {
	return 0x42ff96ed
}

func (*PhoneRequestCallParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) PhoneRequestCall(params *PhoneRequestCallParams) (*PhonePhoneCall, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending PhoneRequestCall")
	}

	resp, ok := responseData.(*PhonePhoneCall)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneSaveCallDebugParams struct {
	Peer  *InputPhoneCall
	Debug *DataJson
}

func (*PhoneSaveCallDebugParams) CRC() uint32 {
	return 0x277add7e
}

// Get the participants of a channel
func (c *Client) PhoneSaveCallDebug(peer *InputPhoneCall, debug *DataJson) (bool, error) {
	responseData, err := c.MakeRequest(&PhoneSaveCallDebugParams{
		Debug: debug,
		Peer:  peer,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending PhoneSaveCallDebug")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneSendSignalingDataParams struct {
	Peer *InputPhoneCall
	Data []byte
}

func (*PhoneSendSignalingDataParams) CRC() uint32 {
	return 0xff7a9383
}

// Get the participants of a channel
func (c *Client) PhoneSendSignalingData(peer *InputPhoneCall, data []byte) (bool, error) {
	responseData, err := c.MakeRequest(&PhoneSendSignalingDataParams{
		Data: data,
		Peer: peer,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending PhoneSendSignalingData")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type PhoneSetCallRatingParams struct {
	UserInitiative bool `tl:"flag:0,encoded_in_bitflags"`
	Peer           *InputPhoneCall
	Rating         int32
	Comment        string
}

func (*PhoneSetCallRatingParams) CRC() uint32 {
	return 0x59ead627
}

func (*PhoneSetCallRatingParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) PhoneSetCallRating(userInitiative bool, peer *InputPhoneCall, rating int32, comment string) (Updates, error) {
	responseData, err := c.MakeRequest(&PhoneSetCallRatingParams{
		Comment:        comment,
		Peer:           peer,
		Rating:         rating,
		UserInitiative: userInitiative,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending PhoneSetCallRating")
	}

	resp, ok := responseData.(Updates)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type PhotosDeletePhotosParams struct {
	ID []InputPhoto
}

func (*PhotosDeletePhotosParams) CRC() uint32 {
	return 0x87cf7f2f
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) PhotosDeletePhotos(id []InputPhoto) ([]int64, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&PhotosDeletePhotosParams{ID: id}, reflect.TypeOf([]int64{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending PhotosDeletePhotos")
	}

	resp, ok := responseData.([]int64)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type PhotosGetUserPhotosParams struct {
	UserID InputUser
	Offset int32
	MaxID  int64
	Limit  int32
}

func (*PhotosGetUserPhotosParams) CRC() uint32 {
	return 0x91cd32a8
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) PhotosGetUserPhotos(userID InputUser, offset int32, maxID int64, limit int32) (PhotosPhotos, error) {
	responseData, err := c.MakeRequest(&PhotosGetUserPhotosParams{
		Limit:  limit,
		MaxID:  maxID,
		Offset: offset,
		UserID: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending PhotosGetUserPhotos")
	}

	resp, ok := responseData.(PhotosPhotos)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type PhotosUpdateProfilePhotoParams struct {
	ID InputPhoto
}

func (*PhotosUpdateProfilePhotoParams) CRC() uint32 {
	return 0x72d4742c
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) PhotosUpdateProfilePhoto(id InputPhoto) (*PhotosPhoto, error) {
	responseData, err := c.MakeRequest(&PhotosUpdateProfilePhotoParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending PhotosUpdateProfilePhoto")
	}

	resp, ok := responseData.(*PhotosPhoto)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type PhotosUploadProfilePhotoParams struct {
	File         InputFile `tl:"flag:0"`
	Video        InputFile `tl:"flag:1"`
	VideoStartTs float64   `tl:"flag:2"`
}

func (*PhotosUploadProfilePhotoParams) CRC() uint32 {
	return 0x89f30f69
}

func (*PhotosUploadProfilePhotoParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) PhotosUploadProfilePhoto(file, video InputFile, videoStartTs float64) (*PhotosPhoto, error) {
	responseData, err := c.MakeRequest(&PhotosUploadProfilePhotoParams{
		File:         file,
		Video:        video,
		VideoStartTs: videoStartTs,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending PhotosUploadProfilePhoto")
	}

	resp, ok := responseData.(*PhotosPhoto)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StatsGetBroadcastStatsParams struct {
	Dark    bool         `tl:"flag:0,encoded_in_bitflags"`
	Channel InputChannel // Channel
}

func (*StatsGetBroadcastStatsParams) CRC() uint32 {
	return 0xab42441a
}

func (*StatsGetBroadcastStatsParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) StatsGetBroadcastStats(dark bool, channel InputChannel) (*StatsBroadcastStats, error) {
	responseData, err := c.MakeRequest(&StatsGetBroadcastStatsParams{
		Channel: channel,
		Dark:    dark,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending StatsGetBroadcastStats")
	}

	resp, ok := responseData.(*StatsBroadcastStats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StatsGetMegagroupStatsParams struct {
	Dark    bool         `tl:"flag:0,encoded_in_bitflags"`
	Channel InputChannel // Channel
}

func (*StatsGetMegagroupStatsParams) CRC() uint32 {
	return 0xdcdf8607
}

func (*StatsGetMegagroupStatsParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) StatsGetMegagroupStats(dark bool, channel InputChannel) (*StatsMegagroupStats, error) {
	responseData, err := c.MakeRequest(&StatsGetMegagroupStatsParams{
		Channel: channel,
		Dark:    dark,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending StatsGetMegagroupStats")
	}

	resp, ok := responseData.(*StatsMegagroupStats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StatsGetMessagePublicForwardsParams struct {
	Channel    InputChannel // Channel
	MsgID      int32        // Message identifier on which a current query depends
	OffsetRate int32
	OffsetPeer InputPeer
	OffsetID   int32
	Limit      int32
}

func (*StatsGetMessagePublicForwardsParams) CRC() uint32 {
	return 0x5630281b
}

// Get the participants of a channel
func (c *Client) StatsGetMessagePublicForwards(params *StatsGetMessagePublicForwardsParams) (MessagesMessages, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending StatsGetMessagePublicForwards")
	}

	resp, ok := responseData.(MessagesMessages)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StatsGetMessageStatsParams struct {
	Dark    bool         `tl:"flag:0,encoded_in_bitflags"`
	Channel InputChannel // Channel
	MsgID   int32        // Message identifier on which a current query depends
}

func (*StatsGetMessageStatsParams) CRC() uint32 {
	return 0xb6e0a3f5
}

func (*StatsGetMessageStatsParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) StatsGetMessageStats(dark bool, channel InputChannel, msgID int32) (*StatsMessageStats, error) {
	responseData, err := c.MakeRequest(&StatsGetMessageStatsParams{
		Channel: channel,
		Dark:    dark,
		MsgID:   msgID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending StatsGetMessageStats")
	}

	resp, ok := responseData.(*StatsMessageStats)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StatsLoadAsyncGraphParams struct {
	Token string
	X     int64 `tl:"flag:0"`
}

func (*StatsLoadAsyncGraphParams) CRC() uint32 {
	return 0x621d5fa0
}

func (*StatsLoadAsyncGraphParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) StatsLoadAsyncGraph(token string, x int64) (StatsGraph, error) {
	responseData, err := c.MakeRequest(&StatsLoadAsyncGraphParams{
		Token: token,
		X:     x,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending StatsLoadAsyncGraph")
	}

	resp, ok := responseData.(StatsGraph)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StickersAddStickerToSetParams struct {
	Stickerset InputStickerSet
	Sticker    *InputStickerSetItem
}

func (*StickersAddStickerToSetParams) CRC() uint32 {
	return 0x8653febe
}

// Get the participants of a channel
func (c *Client) StickersAddStickerToSet(stickerset InputStickerSet, sticker *InputStickerSetItem) (*MessagesStickerSet, error) {
	responseData, err := c.MakeRequest(&StickersAddStickerToSetParams{
		Sticker:    sticker,
		Stickerset: stickerset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending StickersAddStickerToSet")
	}

	resp, ok := responseData.(*MessagesStickerSet)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StickersChangeStickerPositionParams struct {
	Sticker  InputDocument
	Position int32
}

func (*StickersChangeStickerPositionParams) CRC() uint32 {
	return 0xffb6d4ca
}

// Get the participants of a channel
func (c *Client) StickersChangeStickerPosition(sticker InputDocument, position int32) (*MessagesStickerSet, error) {
	responseData, err := c.MakeRequest(&StickersChangeStickerPositionParams{
		Position: position,
		Sticker:  sticker,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending StickersChangeStickerPosition")
	}

	resp, ok := responseData.(*MessagesStickerSet)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StickersCreateStickerSetParams struct {
	Masks     bool `tl:"flag:0,encoded_in_bitflags"`
	Animated  bool `tl:"flag:1,encoded_in_bitflags"`
	UserID    InputUser
	Title     string
	ShortName string
	Thumb     InputDocument `tl:"flag:2"`
	Stickers  []*InputStickerSetItem
}

func (*StickersCreateStickerSetParams) CRC() uint32 {
	return 0xf1036780
}

func (*StickersCreateStickerSetParams) FlagIndex() int {
	return 0
}

// Get the participants of a channel
func (c *Client) StickersCreateStickerSet(params *StickersCreateStickerSetParams) (*MessagesStickerSet, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending StickersCreateStickerSet")
	}

	resp, ok := responseData.(*MessagesStickerSet)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StickersRemoveStickerFromSetParams struct {
	Sticker InputDocument
}

func (*StickersRemoveStickerFromSetParams) CRC() uint32 {
	return 0xf7760f51
}

// Get the participants of a channel
func (c *Client) StickersRemoveStickerFromSet(sticker InputDocument) (*MessagesStickerSet, error) {
	responseData, err := c.MakeRequest(&StickersRemoveStickerFromSetParams{Sticker: sticker})
	if err != nil {
		return nil, errors.Wrap(err, "sending StickersRemoveStickerFromSet")
	}

	resp, ok := responseData.(*MessagesStickerSet)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Get the participants of a channel
type StickersSetStickerSetThumbParams struct {
	Stickerset InputStickerSet
	Thumb      InputDocument
}

func (*StickersSetStickerSetThumbParams) CRC() uint32 {
	return 0x9a364e30
}

// Get the participants of a channel
func (c *Client) StickersSetStickerSetThumb(stickerset InputStickerSet, thumb InputDocument) (*MessagesStickerSet, error) {
	responseData, err := c.MakeRequest(&StickersSetStickerSetThumbParams{
		Stickerset: stickerset,
		Thumb:      thumb,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending StickersSetStickerSetThumb")
	}

	resp, ok := responseData.(*MessagesStickerSet)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UpdatesGetChannelDifferenceParams struct {
	Force   bool `tl:"flag:0,encoded_in_bitflags"`
	Channel InputChannel
	Filter  ChannelMessagesFilter
	Pts     int32
	Limit   int32
}

func (*UpdatesGetChannelDifferenceParams) CRC() uint32 {
	return 0x3173d78
}

func (*UpdatesGetChannelDifferenceParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UpdatesGetChannelDifference(params *UpdatesGetChannelDifferenceParams) (UpdatesChannelDifference, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending UpdatesGetChannelDifference")
	}

	resp, ok := responseData.(UpdatesChannelDifference)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UpdatesGetDifferenceParams struct {
	Pts           int32
	PtsTotalLimit int32 `tl:"flag:0"`
	Date          int32
	Qts           int32
}

func (*UpdatesGetDifferenceParams) CRC() uint32 {
	return 0x25939651
}

func (*UpdatesGetDifferenceParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UpdatesGetDifference(pts, ptsTotalLimit, date, qts int32) (UpdatesDifference, error) {
	responseData, err := c.MakeRequest(&UpdatesGetDifferenceParams{
		Date:          date,
		Pts:           pts,
		PtsTotalLimit: ptsTotalLimit,
		Qts:           qts,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending UpdatesGetDifference")
	}

	resp, ok := responseData.(UpdatesDifference)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UpdatesGetStateParams struct{}

func (*UpdatesGetStateParams) CRC() uint32 {
	return 0xedd4882a
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UpdatesGetState() (*UpdatesState, error) {
	responseData, err := c.MakeRequest(&UpdatesGetStateParams{})
	if err != nil {
		return nil, errors.Wrap(err, "sending UpdatesGetState")
	}

	resp, ok := responseData.(*UpdatesState)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadGetCdnFileParams struct {
	FileToken []byte
	Offset    int32
	Limit     int32
}

func (*UploadGetCdnFileParams) CRC() uint32 {
	return 0x2000bcc3
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadGetCdnFile(fileToken []byte, offset, limit int32) (UploadCdnFile, error) {
	responseData, err := c.MakeRequest(&UploadGetCdnFileParams{
		FileToken: fileToken,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending UploadGetCdnFile")
	}

	resp, ok := responseData.(UploadCdnFile)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadGetCdnFileHashesParams struct {
	FileToken []byte
	Offset    int32
}

func (*UploadGetCdnFileHashesParams) CRC() uint32 {
	return 0x4da54231
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadGetCdnFileHashes(fileToken []byte, offset int32) ([]*FileHash, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&UploadGetCdnFileHashesParams{
		FileToken: fileToken,
		Offset:    offset,
	}, reflect.TypeOf([]*FileHash{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending UploadGetCdnFileHashes")
	}

	resp, ok := responseData.([]*FileHash)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadGetFileParams struct {
	Precise      bool `tl:"flag:0,encoded_in_bitflags"`
	CdnSupported bool `tl:"flag:1,encoded_in_bitflags"`
	Location     InputFileLocation
	Offset       int32
	Limit        int32
}

func (*UploadGetFileParams) CRC() uint32 {
	return 0xb15a9afc
}

func (*UploadGetFileParams) FlagIndex() int {
	return 0
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadGetFile(params *UploadGetFileParams) (UploadFile, error) {
	responseData, err := c.MakeRequest(params)
	if err != nil {
		return nil, errors.Wrap(err, "sending UploadGetFile")
	}

	resp, ok := responseData.(UploadFile)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadGetFileHashesParams struct {
	Location InputFileLocation
	Offset   int32
}

func (*UploadGetFileHashesParams) CRC() uint32 {
	return 0xc7025931
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadGetFileHashes(location InputFileLocation, offset int32) ([]*FileHash, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&UploadGetFileHashesParams{
		Location: location,
		Offset:   offset,
	}, reflect.TypeOf([]*FileHash{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending UploadGetFileHashes")
	}

	resp, ok := responseData.([]*FileHash)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadGetWebFileParams struct {
	Location InputWebFileLocation
	Offset   int32
	Limit    int32
}

func (*UploadGetWebFileParams) CRC() uint32 {
	return 0x24e6818d
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadGetWebFile(location InputWebFileLocation, offset, limit int32) (*UploadWebFile, error) {
	responseData, err := c.MakeRequest(&UploadGetWebFileParams{
		Limit:    limit,
		Location: location,
		Offset:   offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "sending UploadGetWebFile")
	}

	resp, ok := responseData.(*UploadWebFile)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadReuploadCdnFileParams struct {
	FileToken    []byte
	RequestToken []byte
}

func (*UploadReuploadCdnFileParams) CRC() uint32 {
	return 0x9b2754a8
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadReuploadCdnFile(fileToken, requestToken []byte) ([]*FileHash, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&UploadReuploadCdnFileParams{
		FileToken:    fileToken,
		RequestToken: requestToken,
	}, reflect.TypeOf([]*FileHash{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending UploadReuploadCdnFile")
	}

	resp, ok := responseData.([]*FileHash)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadSaveBigFilePartParams struct {
	FileID         int64
	FilePart       int32
	FileTotalParts int32
	Bytes          []byte
}

func (*UploadSaveBigFilePartParams) CRC() uint32 {
	return 0xde7b673d
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadSaveBigFilePart(fileID int64, filePart, fileTotalParts int32, bytes []byte) (bool, error) {
	responseData, err := c.MakeRequest(&UploadSaveBigFilePartParams{
		Bytes:          bytes,
		FileID:         fileID,
		FilePart:       filePart,
		FileTotalParts: fileTotalParts,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending UploadSaveBigFilePart")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UploadSaveFilePartParams struct {
	FileID   int64
	FilePart int32
	Bytes    []byte
}

func (*UploadSaveFilePartParams) CRC() uint32 {
	return 0xb304a621
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UploadSaveFilePart(fileID int64, filePart int32, bytes []byte) (bool, error) {
	responseData, err := c.MakeRequest(&UploadSaveFilePartParams{
		Bytes:    bytes,
		FileID:   fileID,
		FilePart: filePart,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending UploadSaveFilePart")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UsersGetFullUserParams struct {
	ID InputUser
}

func (*UsersGetFullUserParams) CRC() uint32 {
	return 0xca30a5b1
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UsersGetFullUser(id InputUser) (*UserFull, error) {
	responseData, err := c.MakeRequest(&UsersGetFullUserParams{ID: id})
	if err != nil {
		return nil, errors.Wrap(err, "sending UsersGetFullUser")
	}

	resp, ok := responseData.(*UserFull)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UsersGetUsersParams struct {
	ID []InputUser
}

func (*UsersGetUsersParams) CRC() uint32 {
	return 0xd91a548
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UsersGetUsers(id []InputUser) ([]User, error) {
	responseData, err := c.MakeRequestWithHintToDecoder(&UsersGetUsersParams{ID: id}, reflect.TypeOf([]User{}))
	if err != nil {
		return nil, errors.Wrap(err, "sending UsersGetUsers")
	}

	resp, ok := responseData.([]User)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
type UsersSetSecureValueErrorsParams struct {
	ID     InputUser
	Errors []SecureValueError
}

func (*UsersSetSecureValueErrorsParams) CRC() uint32 {
	return 0x90c894b5
}

// Sends a Telegram Passport authorization form, effectively sharing data with the service
func (c *Client) UsersSetSecureValueErrors(id InputUser, errs []SecureValueError) (bool, error) {
	responseData, err := c.MakeRequest(&UsersSetSecureValueErrorsParams{
		Errors: errs,
		ID:     id,
	})
	if err != nil {
		return false, errors.Wrap(err, "sending UsersSetSecureValueErrors")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}
