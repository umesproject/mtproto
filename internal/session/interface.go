package session

// Sesion is a basic data of specific session. Typically, session stores default hostname of mtproto server
// (cause all accounts ties to specific server after sign in), session key, server hash and salt.
type Session struct {
	Key      []byte
	Hash     []byte
	Salt     int64
	Hostname string
}
