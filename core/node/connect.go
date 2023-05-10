package node

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func Connect(node *host.Host, address string) (*peer.AddrInfo, error) {
	addr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return nil, err
	}
	new, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return nil, err
	}
	return new, (*node).Connect(context.Background(), *new)
}
