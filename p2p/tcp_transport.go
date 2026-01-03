package p2p

import (
	"io"
	"log"
	"net"
)

var (
	TransportProtocol = "tcp"
)

type TCPPeer struct {
	conn     net.Conn // represents the connection
	outbound bool
	// if a node/peer dialUp to other then it would be true
	// also help to differential the handleNewConnection to differentiate
	// whenever it has to handle the new connection
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) RemoteAddress() string {
	return p.conn.RemoteAddr().String()
}

type TCPTransportOpts struct {
	ListenAddress string
	HandshakeFunc HandshakeFunc
	Decode        Decoder
	OnPeer        OnPeer
}

type TCPTransport struct {
	listener net.Listener
	TCPTransportOpts
	rpchan chan RPC
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpchan:           make(chan RPC), //unbuffered channel: A send blocks until some goroutine is receiving.
	}
}

// ListenAndAccept() implements the Transport interface to start listening
// and accepting new connections
func (t *TCPTransport) ListenAndAccept() error {

	// 1. Server listen on provided address and transport protocol
	var err error
	t.listener, err = net.Listen(TransportProtocol, t.ListenAddress)
	if err != nil {
		return err
	}
	log.Printf("Server started listening: %+v\n", t.ListenAddress)

	// 2. Server accept the new connection on its listening address
	// make it indepdent by goroutine,
	//  otherwise it has loop, which will block ListenAndAccept
	// Now it seperate (but depend on listener), so it only should s
	// top when server listener died
	go t.AcceptLoop()

	return nil
}

func (t *TCPTransport) AcceptLoop() {

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
		go t.handleNewConnection(conn, false)
	}
}

func (t *TCPTransport) handleNewConnection(conn net.Conn, outbound bool) {
	log.Printf("Handling the upcoming connection: %+v\n", conn)
	defer conn.Close()

	// Create a Peer (in simple terms its a conn)
	peer := NewTCPPeer(conn, outbound)

	// Handshake
	if t.HandshakeFunc != nil { // To Check if the func defined or not
		if err := t.HandshakeFunc(peer); err != nil {
			log.Printf("handshake failed:%+v\n", err)
			return
		}
	}

	// OnPeer Action
	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			log.Printf("onpeer action err: %+v\n", err)
			return
		}
	}

	msg := RPC{}

	for {

		// Decode the message on the connection stream - have blocking call
		if err := t.Decode.Decode(conn, &msg); err != nil {
			log.Printf("message decode failed:%+v\n", err)
			// stop reading in case of connection break or closed
			// not in the case when some func related errors happens
			if err == io.EOF || err == net.ErrClosed {
				return
			}
		}
		// After Read() from Decode reading the stream of conn
		msg.From = conn.RemoteAddr().String()
		// log.Printf("recieved msg %v from %v\n", string(msg.Payload), msg.From)
		t.rpchan <- msg // send rpc message to channel (can be fetch by other connections)

	}

}

func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpchan
}

func (t *TCPTransport) Dial(address string) error {
	_, err := net.Dial(TransportProtocol, address)
	return err
}
