// AIP2P Core package. Implement node, logs and p2p services
package core

import (
	"fmt"

	"webimizer.dev/aip2p/core/node"
)

// AIP2P application entry function
func StartApplication() {
	fmt.Printf("AIP2P Application %s build %d\n\n", VERSION, BUILD_NUMBER)
	node.StartNode()
}
