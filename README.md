# learnforeverstore

<details>
<summary> PHASE - I - TCP Implementation</summary>

- 1.1: Simple TCP prototype defined 
- 1.2: Handshake functionality addded using callback func
- 1.3: handleNewConnection modifed to read message of its connection (send from telnet)
- 1.3: (Revised) Refactor
- 1.4: Decoder and Refactor (added TCPTransportOpts)
- 1.5: RPC Message implemented (to define payload for comms), to also take the message out from Decoder func to TCPTransport
- 1.6: Peer has been defined with TCPPeer and some refactor (replacing conn with peer)
- 1.7 - OnPeer func added - to take action or can be used to as Notification to take action when the connection (peer) is establishedv

</details>

> **Note:**
> At the point, we are able to create the TCP lib, to which following tasks
> can be done
> - Server(Node) instance start to listen on a sepcified port
> - Upcoming new connections after being accepted, will be handled independently (Handshake, Peer Conversion, Decode) and reading from its conn READ stream to stdout RPC Message
> - Similarly that independent connection will be closed on disconnection from Server (cause EOF) but except this error, it will keep reading unless Server is down

<details>
<summary> PHASE - II - Store Implementation</summary>

- 2.1: Store created with writeStream, to create a file on the disk

</details>