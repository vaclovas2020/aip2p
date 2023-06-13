// Implements AIP2P node based on libp2p
package node

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"webimizer.dev/aip2p/core/protocol/dnn"
)

// Listen on tcp/ip4 on random port
func Listen(logInfoHandler dnn.LogInfoFunc, logErrorHandler dnn.LogErrorFunc) (*host.Host, error) {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/0",
		),
	)
	if err != nil {
		return nil, err
	}
	dnnService, err := dnn.NewDnnService(&node, logInfoHandler, logErrorHandler)
	node.SetStreamHandler(dnn.ID, dnnService.StreamHandler)
	return &node, err
}
