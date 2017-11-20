package main

import (
	"github.com/mojachieee/go-HoneyPot/config"
	"github.com/mojachieee/go-HoneyPot/database"
	"github.com/mojachieee/go-HoneyPot/tcp"
)

func main() {
	config := config.Read()
	db := database.InitDatabase(config.DB)
	defer db.Close()
	tcpServer := tcp.NewServer(config.TCP.Ports)
	tcpServer.Start(db, config.DB)
}
