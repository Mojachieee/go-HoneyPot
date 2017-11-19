package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

// Server is the tcp server struct
type Server struct {
	Ports []string
}

// NewServer creates a new tcp server
func NewServer() *Server {
	ports := []string{"8992", "1280"}
	return &Server{ports}
}

// Start starts the tcp server
func (t *Server) Start(db *gorm.DB) {
	var wg sync.WaitGroup
	wg.Add(len(t.Ports))
	for _, port := range t.Ports {
		go func(port string, wg *sync.WaitGroup, db *gorm.DB) {
			fmt.Printf("Listening on tcp port: %v\n", port)
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
				go handleConnection(conn, db)
			}
		}(port, &wg, db)
	}
	wg.Wait()
	fmt.Println("TCP Server Stopped")
}

func handleConnection(conn net.Conn, db *gorm.DB) {
	fmt.Println("connection")
	data := make([]byte, 4096)
	n, err := conn.Read(data)
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}
	defer conn.Close()
	fmt.Printf("Received data from %v, of length %v data is %v\n", conn.RemoteAddr(), n, data[:n])
	remHost, remPort, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		fmt.Printf("Failed to split remote host and port: %v\n", err)
		return
	}
	locHost, locPort, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		fmt.Printf("Failed to split remote host and port: %v\n", err)
		return
	}
	str := fmt.Sprintf(`INSERT INTO honeyinfo (Date, InIp, InPort, DestIP, DestPort, SessionID, DataLength)VALUES ("%v", "%v", "%v", "%v", "%v", "%v", "%v")`,
		time.Now().Format("20060102150405"), remHost, remPort, locHost, locPort, 0, n)
	db.Exec(str)
}
