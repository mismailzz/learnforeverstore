package p2p

import "net"

// Callback func. -> to this type handshake func can be
// injected as dependency injection
// HandshakeFunc can be injected - incase any any logic required for the conn
type HandshakeFunc func(net.Conn) error

func NoHandshakeFunc(conn net.Conn) error {
	return nil
}
