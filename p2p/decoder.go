package p2p

import (
	"log"
	"net"
)

// Different type of Decoder can be used to decode the messages
// being sent on the communication channel of p2p network
type Decoder interface {
	Decode(net.Conn) error
}

type DefaultDecoder struct{}

func (d *DefaultDecoder) Decode(conn net.Conn) error {

	buff := make([]byte, 1024)
	// blocking read call
	_, err := conn.Read(buff)
	if err != nil {
		log.Printf("read error: %v from %v\n", err, conn)
		return err
	}
	log.Printf("recieved msg: %v\n", string(buff))

	return nil
}
