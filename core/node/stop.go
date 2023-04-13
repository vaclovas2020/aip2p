package node

import (
	"github.com/libp2p/go-libp2p/core/host"
)

func StopListen(node *host.Host) {
	if node == nil {
		return
	}
	if err := (*node).Close(); err != nil {
		panic(err)
	}
}
