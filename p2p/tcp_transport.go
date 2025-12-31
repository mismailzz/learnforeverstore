package p2p

import (
	"io"
	"log"
	"net"
)

type TCPTransportOpts struct {
	ListenAddress string
	HandshakeFunc HandshakeFunc
	Decode        Decoder
}

type TCPTransport struct {
	listener net.Listener
	TCPTransportOpts
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

// ListenAndAccept() implements the Transport interface to start listening
// and accepting new connections
func (t *TCPTransport) ListenAndAccept() error {

	// 1. Server listen on provided address and transport protocol
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	log.Printf("Server started listening: %+v\n", t.ListenAddress)

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

	// Handshake
	if t.HandshakeFunc != nil { // To Check if the func defined or not
		if err := t.HandshakeFunc(conn); err != nil {
			log.Printf("handshake failed:%+v\n", err)
			return
		}
	}

	for {

		// Decode the message on the connection stream
		if err := t.Decode.Decode(conn); err != nil {
			log.Printf("message decode failed:%+v\n", err)
			// stop reading in case of connection break or closed
			// not in the case when some func related errors happens
			if err == io.EOF || err == net.ErrClosed {
				return
			}
		}

	}

}
