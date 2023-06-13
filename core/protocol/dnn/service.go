package dnn

import (
	"bufio"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
)

type Aip2pService struct {
	Host *host.Host
}

const ID string = "/dnn/1.0.0"

func NewAip2pService(h *host.Host) (*Aip2pService, error) {
	return &Aip2pService{Host: h}, nil
}

func (s *Aip2pService) StreamHandler(buff network.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(buff), bufio.NewWriter(buff))

	go readData(rw)
	go writeData(rw, []byte("Hello World\n"))

	// stream 's' will stay open until you close it (or the other side closes it).
}

func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received: %s", str)
	}
}

func writeData(rw *bufio.ReadWriter, p []byte) {
	rw.Write(p)
	rw.Flush()
}
