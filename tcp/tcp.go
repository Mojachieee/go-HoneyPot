package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// Server is the tcp server struct
type Server struct {
	Ports []string
}

// NewServer creates a new tcp server
func NewServer() *Server {
	ports := []string{"80", "8080", "1", "2"}
	return &Server{ports}
}

// Start starts the tcp honeypot
func (t *Server) Start() {
	var wg sync.WaitGroup
	wg.Add(len(t.Ports))
	for _, port := range t.Ports {
		go func(port string, wg *sync.WaitGroup) {
			fmt.Printf("Listening on port: %v\n", port)
			listen, err := net.Listen("tcp", ":"+port)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			for {
				conn, err := listen.Accept()
				if err != nil {
					log.Fatal(err)
					// handle error
				}
				go handleConnection(conn)
			}
		}(port, &wg)
	}
	wg.Wait()
	fmt.Println("TCP Server Stopped")
}

func handleConnection(con net.Conn) {
	fmt.Println("connection")
}
