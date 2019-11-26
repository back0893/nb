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
	server.AddRouter(0x1001, router.NewHandler())
	server.Run()
}
