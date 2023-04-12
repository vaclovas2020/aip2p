package node

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p/core/host"
	"webimizer.dev/aip2p/core/logs"
)

func ReceiveExitSignal(node host.Host) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logs.LogWarning("Received signal, shutting down...", "aip2p")

	// shut the node down
	if err := node.Close(); err != nil {
		logs.LogError(err, "aip2p", true)
	}
}
