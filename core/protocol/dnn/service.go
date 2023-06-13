package dnn

import (
	"bufio"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

type DnnService struct {
	Host                 *host.Host
	LogInfoHandler       LogInfoFunc
	LogErrorHandler      LogErrorFunc
	AddPeerToListHandler AddPeerToListFunc
}

type LogInfoFunc func(string, ...interface{})
type LogErrorFunc func(error)
type AddPeerToListFunc func(*peer.AddrInfo)

const ID protocol.ID = "/dnn/1.0.0"

func NewDnnService(h *host.Host, logInfoHandler LogInfoFunc, logErrorHandler LogErrorFunc, addPeerToListHandler AddPeerToListFunc) (*DnnService, error) {
	return &DnnService{Host: h, LogInfoHandler: logInfoHandler, LogErrorHandler: logErrorHandler, AddPeerToListHandler: addPeerToListHandler}, nil
}

func (s *DnnService) StreamHandler(buff network.Stream) {
	addrs := make([]multiaddr.Multiaddr, 0, 1)
	addrs = append(addrs, buff.Conn().RemoteMultiaddr())
	s.AddPeerToListHandler(&peer.AddrInfo{ID: buff.Conn().RemotePeer(), Addrs: addrs})
	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(buff), bufio.NewWriter(buff))

	go s.readData(rw, buff)
	go s.writeData(rw, []byte("Hello World\n"), buff)

	// stream 's' will stay open until you close it (or the other side closes it).
}

func (s *DnnService) readData(rw *bufio.ReadWriter, buff network.Stream) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			s.LogErrorHandler(err)
			return
		}
		s.LogInfoHandler("Received from %s: %s", buff.Conn().RemoteMultiaddr().String(), str)
	}
}

func (s *DnnService) writeData(rw *bufio.ReadWriter, p []byte, buff network.Stream) {
	s.LogInfoHandler("Sending %d bytes to %s...", len(p), buff.Conn().RemoteMultiaddr().String())
	rw.Write(p)
	rw.Flush()
}
