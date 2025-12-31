package p2p

// RPC is the message which is being transmitted
// our the p2p network communication
type RPC struct {
	From    string
	Payload []byte
}
