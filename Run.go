package main

import (
	"Nb/net"
	"Nb/protocol"
	"Nb/router"
)

func main() {
	server := net.NewServer()
	jt809 := &protocol.Jt809{}
	server.SetProtocol(jt809)
	server.AddRouter(0x1001, router.NewJt1001())
	server.AddRouter(0x1200, router.NewJt1200())
	server.AddRouter(0x1005, router.NewJt1005())
	server.Run()
}
