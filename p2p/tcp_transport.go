package p2p

import (
	"errors"
	"io"
	"log"
	"net"
)

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

// ListenAndAccept() implements the Transport interface to start listening
// and accepting new connections
func (t *TCPTransport) ListenAndAccept() error {

	// 1. Server listen on provided address and transport protocol
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
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

	defer conn.Close()

	if t.handshakeFunc != nil { // To Check if the func defined or not
		if err := t.handshakeFunc(conn); err != nil {
			log.Printf("handshake failed:%+v\n", err)
			return
		}
	}

	for {
		buff := make([]byte, 1024)
		// blocking read call
		_, err := conn.Read(buff)
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Printf("connection already closed: %v\n", conn)
				return
			}
			if err == io.EOF {
				log.Printf("client closed connection: %v\n", conn)
				return
			}
			log.Printf("read error: %v from %v\n", err, conn)
			return
		}

	}

}
