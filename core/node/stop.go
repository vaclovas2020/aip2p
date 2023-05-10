package node

import (
	"github.com/libp2p/go-libp2p/core/host"
)

func StopListen(node *host.Host) error {
	if node == nil {
		return nil
	}
	return (*node).Close()
}
