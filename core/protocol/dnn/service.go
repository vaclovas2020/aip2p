package dnn

import (
	"bufio"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type DnnService struct {
	Host *host.Host
}

const ID protocol.ID = "/dnn/1.0.0"

func NewDnnService(h *host.Host) (*DnnService, error) {
	return &DnnService{Host: h}, nil
}

func (s *DnnService) StreamHandler(buff network.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(buff), bufio.NewWriter(buff))

	go s.readData(rw)
	go s.writeData(rw, []byte("Hello World\n"))

	// stream 's' will stay open until you close it (or the other side closes it).
}

func (s *DnnService) readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received from %s: %s", (*s.Host).ID().String(), str)
	}
}

func (s *DnnService) writeData(rw *bufio.ReadWriter, p []byte) {
	rw.Write(p)
	rw.Flush()
}
