package p2p

import (
	"log"
	"net"
)

const TransportProtocol = "tcp"

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshakeFunc HandshakeFunc
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	// 1. Server listen on provided address and transport protocol
	var err error
	t.listener, err = net.Listen(TransportProtocol, t.listenAddress)
	if err != nil {
		return err
	}
	log.Printf("Server started listening: %+v\n", t.listenAddress)

	// 2. Server accept the new connection on its listening address
	t.acceptLoop()

	return nil
}

func (t *TCPTransport) acceptLoop() {

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Printf("accept error:%+v\n", err)
			continue
			// No need to halt the listener
			// can proceed for other connection Accept requests
			// also failed one can be supposidly be tried from client side
		}

		// As connetion accepted, then proceed to handle that connection
		// seperate and independent
		go t.handleNewConnection(conn)
	}
}

func (t *TCPTransport) handleNewConnection(conn net.Conn) {
	log.Printf("Handling the upcoming connection: %+v\n", conn)

	if t.handshakeFunc != nil { // To Check if the func defined or not
		if err := t.handshakeFunc(conn); err != nil {
			log.Printf("handshake failed:%+v\n", err)
			return
		}
	}

}
