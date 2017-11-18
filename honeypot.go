package main

import "github.com/mojachieee/go-honeypot/tcp"

func main() {
	tcpServer := tcp.NewServer()
	tcpServer.Start()
}
