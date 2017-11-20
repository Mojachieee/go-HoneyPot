package main

import (
	"github.com/mojachieee/go-HoneyPot/config"
	"github.com/mojachieee/go-HoneyPot/database"
	"github.com/mojachieee/go-HoneyPot/tcp"
)

func main() {
	cfg := config.Read()
	db := database.InitDatabase(cfg.DB)
	defer db.Close()
	tcpServer := tcp.NewServer(cfg.TCP.Ports)
	tcpServer.Start(db, cfg.DB)
}
