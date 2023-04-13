package node

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p/core/host"
)

func ReceiveExitSignal(node host.Host) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}
}
