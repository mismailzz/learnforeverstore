package p2p

// Callback func. -> to this type handshake func can be
// injected as dependency injection
// HandshakeFunc can be injected - incase any any logic required for the conn
type HandshakeFunc func(Peer) error

func NoHandshakeFunc(peer Peer) error {
	return nil
}
