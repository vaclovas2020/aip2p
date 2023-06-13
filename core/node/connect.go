package node

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/multiformats/go-multiaddr"
	"webimizer.dev/aip2p/core/protocol/dnn"
)

func Connect(node *host.Host, address string, logInfoHandler dnn.LogInfoFunc, logErrorHandler dnn.LogErrorFunc) (*peer.AddrInfo, error) {
	addr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return nil, err
	}
	new, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return nil, err
	}
	dnnService, err := dnn.NewDnnService(node, logInfoHandler, logErrorHandler)
	if err != nil {
		return nil, err
	}
	s, err := (*node).NewStream(context.Background(), (*new).ID, dnn.ID)
	if err != nil {
		return nil, err
	}
	(*node).Peerstore().AddAddrs((*new).ID, (*new).Addrs, peerstore.PermanentAddrTTL)
	dnnService.StreamHandler(s)
	return new, nil
}
