package node

import (
	"github.com/libp2p/go-libp2p/core/host"
	"webimizer.dev/aip2p/core/protocol/dnn"
)

func Connect(
	node *host.Host, address string,
	logInfoHandler dnn.LogInfoFunc,
	logErrorHandler dnn.LogErrorFunc,
	addPeerToListHandler dnn.AddPeerToListFunc,
	removePeerFromListHandler dnn.RemovePeerFromListFunc,
) error {
	return dnn.Connect(node, address, logInfoHandler, logErrorHandler, addPeerToListHandler, removePeerFromListHandler)
}
