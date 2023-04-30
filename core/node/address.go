package node

import (
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func GetAddress(node *host.Host) ([]multiaddr.Multiaddr, error) {
	peerInfo := peer.AddrInfo{
		ID:    (*node).ID(),
		Addrs: (*node).Addrs(),
	}
	return peer.AddrInfoToP2pAddrs(&peerInfo)
}
