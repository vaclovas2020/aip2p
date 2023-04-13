// Implements AIP2P node based on libp2p
package node

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

// Listen on tcp/ip4 and tcp/ip6 on port 7777
func Listen() (*host.Host, error) {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/7777",
			"/ip6/::/tcp/7777",
		),
	)
	return &node, err
}
