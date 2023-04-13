package aip2p

import (
	"fmt"

	"webimizer.dev/aip2p/core"
	"webimizer.dev/aip2p/core/node"
)

// AIP2P application entry function
func StartApplication() {
	fmt.Printf("AIP2P Application %s build %d\n\n", core.VERSION, core.BUILD_NUMBER)
	node.StartNode()
}
