package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      Netaddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[Netaddr]*LocalTransport
}

func NewLocalTransport(addr Netaddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[Netaddr]*LocalTransport),
	}
}
func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}
func (t *LocalTransport) Connect(tr *LocalTransport) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.peers[tr.Addr()] = tr
	return nil
}
func (t *LocalTransport) SendMessage(to Netaddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()
	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}
	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}
	return nil
}

func (t *LocalTransport) Addr() Netaddr {
	return t.addr
}
