package dnn

import (
	"bufio"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type DnnService struct {
	Host            *host.Host
	LogInfoHandler  LogInfoFunc
	LogErrorHandler LogErrorFunc
}

type LogInfoFunc func(string, ...interface{})
type LogErrorFunc func(error)

const ID protocol.ID = "/dnn/1.0.0"

func NewDnnService(h *host.Host, logInfoHandler LogInfoFunc, logErrorHandler LogErrorFunc) (*DnnService, error) {
	return &DnnService{Host: h, LogInfoHandler: logInfoHandler, LogErrorHandler: logErrorHandler}, nil
}

func (s *DnnService) StreamHandler(buff network.Stream) {
	s.LogInfoHandler("Got a new stream!")
	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(buff), bufio.NewWriter(buff))

	go s.readData(rw, buff)
	go s.writeData(rw, []byte("Hello World\n"))

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

func (s *DnnService) writeData(rw *bufio.ReadWriter, p []byte) {
	rw.Write(p)
	rw.Flush()
}
