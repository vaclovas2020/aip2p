package dnn

import (
	"github.com/libp2p/go-libp2p/core/peer"
)

type NodeList struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Address string  `json:"address"`
	ID      peer.ID `json:"id"`
}
