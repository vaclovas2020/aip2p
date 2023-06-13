package node

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"webimizer.dev/aip2p/core/protocol/dnn"
)

func Connect(node *host.Host, address string, logInfoHandler dnn.LogInfoFunc, logErrorHandler dnn.LogErrorFunc, addPeerToListHandler dnn.AddPeerToListFunc) error {
	addr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return err
	}
	new, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return err
	}
	dnnService, err := dnn.NewDnnService(node, logInfoHandler, logErrorHandler, addPeerToListHandler)
	if err != nil {
		return err
	}
	err = (*node).Connect(context.Background(), *new)
	if err != nil {
		return err
	}
	s, err := (*node).NewStream(context.Background(), (*new).ID, dnn.ID)
	if err != nil {
		return err
	}
	dnnService.StreamHandler(s)
	return nil
}
