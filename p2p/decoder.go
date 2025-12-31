package p2p

import (
	"log"
	"net"
)

// Different type of Decoder can be used to decode the messages
// being sent on the communication channel of p2p network
type Decoder interface {
	Decode(net.Conn, *RPC) error
}

type DefaultDecoder struct{}

func (d *DefaultDecoder) Decode(conn net.Conn, rpc *RPC) error {

	buff := make([]byte, 1024)
	// blocking read call
	n, err := conn.Read(buff) // n is the number of bytes actually read
	if err != nil {
		log.Printf("read error: %v from %v\n", err, conn)
		return err
	}

	rpc.Payload = buff[:n]
	// buff[:n] instead of buff because Read() call doesnt full the capacity,
	// so we can specify the actual bytes read

	return nil
}
