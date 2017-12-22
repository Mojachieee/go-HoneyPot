package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mojachieee/go-HoneyPot/config"
)

// Server is the tcp server struct
type Server struct {
	Ports []string
}

// NewServer creates a new tcp server
func NewServer(ports []string) *Server {
	return &Server{ports}
}

// Start starts the tcp server
func (t *Server) Start(db *gorm.DB, cfg config.Database) {
	var wg sync.WaitGroup
	wg.Add(len(t.Ports))
	for _, port := range t.Ports {
		go func(port string, wg *sync.WaitGroup, db *gorm.DB, cfg config.Database) {
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
				go handleConnection(conn, db, cfg)
			}
		}(port, &wg, db, cfg)
	}
	wg.Wait()
	fmt.Println("TCP Server Stopped")
}

func handleConnection(conn net.Conn, db *gorm.DB, cfg config.Database) {
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
	str := fmt.Sprintf(`INSERT INTO %v (Date, InIp, InPort, DestIP, DestPort, DataLength)VALUES ("%v", "%v", "%v", "%v", "%v", "%v")`,
		cfg.Table, time.Now().Format("20060102150405"), remHost, remPort, locHost, locPort, n)
	db.Exec(str)
}
