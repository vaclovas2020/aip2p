// Implements AIP2P node based on libp2p
package node

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

// Listen on tcp/ip4 on random port
func Listen() (*host.Host, error) {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/0",
		),
	)
	return &node, err
}
