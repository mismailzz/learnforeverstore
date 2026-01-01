package main

import "bytes"

func main() {

	// opts := p2p.TCPTransportOpts{
	// 	ListenAddress: ":3000",
	// 	HandshakeFunc: p2p.NoHandshakeFunc,
	// 	Decode:        &p2p.DefaultDecoder{},
	// }
	// t := p2p.NewTCPTransport(opts)
	// if err := t.ListenAndAccept(); err != nil {
	// 	log.Fatal(err)
	// }

	opts := StoreOpts{
		rootDir: "db",
	}

	store := NewStore(opts)
	store.writeStream("example.txt", bytes.NewReader([]byte("hello world")))

}
