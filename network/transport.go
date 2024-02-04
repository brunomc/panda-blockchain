package network

type Netaddr string
type RPC struct {
	From    Netaddr
	Payload []byte
}
type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(Netaddr, []byte) error
	Addr() Netaddr
}
