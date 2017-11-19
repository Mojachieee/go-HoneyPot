package udp

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// Server is the udp server struct
type Server struct {
	Ports []string
}

// NewServer creates a new udp server
func NewServer() *Server {
	ports := []string{"80", "8080", "1", "2"}
	return &Server{ports}
}

// Start starts the udp server
func (t *Server) Start() {
	var wg sync.WaitGroup
	wg.Add(len(t.Ports))
	for _, port := range t.Ports {
		go func(port string, wg *sync.WaitGroup) {
			fmt.Printf("Listening on udp port: %v\n", port)
			addr, err := net.ResolveUDPAddr("udp", ":"+port)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			conn, err := net.ListenUDP("udp", addr)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			for {
				handleConnection(conn)
			}
		}(port, &wg)
	}
	wg.Wait()
	fmt.Println("UDP Server Stopped")
}

func handleConnection(con net.Conn) {
}
