package dnn

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

type DnnService struct {
	Host                      *host.Host
	LogInfoHandler            LogInfoFunc
	LogErrorHandler           LogErrorFunc
	AddPeerToListHandler      AddPeerToListFunc
	RemovePeerFromListHandler RemovePeerFromListFunc
}

type LogInfoFunc func(string, ...interface{})
type LogErrorFunc func(error)
type AddPeerToListFunc func(*peer.AddrInfo)
type RemovePeerFromListFunc func(peer.ID)

const ID protocol.ID = "/dnn/1.0.0"

func NewDnnService(
	h *host.Host,
	logInfoHandler LogInfoFunc,
	logErrorHandler LogErrorFunc,
	addPeerToListHandler AddPeerToListFunc,
	removePeerFromListHandler RemovePeerFromListFunc,
) (*DnnService, error) {
	return &DnnService{
		Host:                      h,
		LogInfoHandler:            logInfoHandler,
		LogErrorHandler:           logErrorHandler,
		AddPeerToListHandler:      addPeerToListHandler,
		RemovePeerFromListHandler: removePeerFromListHandler,
	}, nil
}

func (s *DnnService) StreamHandler(buff network.Stream) {
	go func() {
		addrs := make([]multiaddr.Multiaddr, 0, 1)
		addrs = append(addrs, buff.Conn().RemoteMultiaddr())
		s.AddPeerToListHandler(&peer.AddrInfo{ID: buff.Conn().RemotePeer(), Addrs: addrs})
		// Create a buffer stream for non blocking read and write.
		rw := bufio.NewReadWriter(bufio.NewReader(buff), bufio.NewWriter(buff))

		// Create input and output channels to read and write data.
		input := make(chan *Message)
		output := make(chan *Message)

		// Create a thread to read and write data.
		go s.readData(rw, buff, input)
		go s.writeData(rw, output, buff)
		go s.receiveMessage(buff, input)

		list := &Message{Type: TYPE_NODELIST, Data: s.prepareNodeList(buff)}
		output <- list
	}()
}

func Connect(
	node *host.Host, address string,
	logInfoHandler LogInfoFunc,
	logErrorHandler LogErrorFunc,
	addPeerToListHandler AddPeerToListFunc,
	removePeerFromListHandler RemovePeerFromListFunc,
) error {
	addr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return err
	}
	new, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return err
	}
	err = (*node).Connect(context.Background(), *new)
	if err != nil {
		return err
	}
	dnnService, err := NewDnnService(node, logInfoHandler, logErrorHandler, addPeerToListHandler, removePeerFromListHandler)
	if err != nil {
		return err
	}
	s, err := (*node).NewStream(context.Background(), (*new).ID, ID)
	if err != nil {
		return err
	}
	dnnService.StreamHandler(s)
	return nil
}

func (s *DnnService) prepareNodeList(buff network.Stream) string {
	nodes := make([]Node, 0)
	for _, node := range (*s.Host).Peerstore().Peers() {
		if node != buff.Conn().RemotePeer() && (*s.Host).ID() != node {
			addrs := (*s.Host).Peerstore().Addrs(node)
			nodes = append(nodes, Node{ID: node, Address: addrs[len(addrs)-1].String()})
		}
	}
	p, err := json.Marshal(NodeList{Nodes: nodes})
	if err != nil {
		s.LogErrorHandler(err)
		return ""
	}
	return string(p)
}

func (s *DnnService) receiveNodeList(buff network.Stream, list *NodeList) {
	for _, p := range list.Nodes {
		if p.ID != buff.Conn().RemotePeer() && (*s.Host).ID() != p.ID {
			addrs := (*s.Host).Peerstore().Addrs(p.ID)
			if len(addrs) == 0 {
				Connect(s.Host, fmt.Sprintf("%s/p2p/%v", p.Address, p.ID), s.LogInfoHandler, s.LogErrorHandler, s.AddPeerToListHandler, s.RemovePeerFromListHandler)
			}
		}
	}
}

func (s *DnnService) receiveMessage(buff network.Stream, msg <-chan *Message) {
	message := <-msg
	switch message.Type {
	case TYPE_NODELIST:
		list := NodeList{}
		err := json.Unmarshal([]byte(message.Data), &list)
		if err != nil {
			s.LogErrorHandler(err)
			return
		}
		s.receiveNodeList(buff, &list)
	default:
		s.LogErrorHandler(errors.New("unknown message type `" + message.Type + "`"))
	}
}

func (s *DnnService) readData(rw *bufio.ReadWriter, buff network.Stream, msg chan<- *Message) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			s.RemovePeerFromListHandler(buff.Conn().RemotePeer())
			return
		}
		s.LogInfoHandler("Received %d bytes from %s: %s", len(str), buff.Conn().RemoteMultiaddr().String(), str)
		message := Message{}
		err = json.Unmarshal([]byte(str), &message)
		if err != nil {
			s.LogErrorHandler(err)
			return
		}
		msg <- &message
	}
}

func (s *DnnService) writeData(rw *bufio.ReadWriter, message <-chan *Message, buff network.Stream) {
	p, err := json.Marshal(<-message)
	if err != nil {
		s.LogErrorHandler(err)
		return
	}
	p = append(p, '\n')
	str := string(p)
	rw.WriteString(str)
	rw.Flush()
	s.LogInfoHandler("Sended %d bytes to %s: %s", len(str), buff.Conn().RemoteMultiaddr().String(), str)
}
