package node

import (
	"fmt"

	"webimizer.dev/aip2p/core/logs"
)

func StartNode() {
	node, err := Listen()
	if err != nil {
		panic(err)
	}

	// print the node's listening addresses and node ID
	logs.Log(fmt.Sprintf("Listen addresses: %v\nNode ID: %v", node.Addrs(), node.ID()), "aip2p")

	// waiting for os.Signal
	ReceiveExitSignal(node)
}
