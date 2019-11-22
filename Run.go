package main

import (
	"Nb/net"
	"Nb/router"
)

func main() {
	server := net.NewServer()
	server.AddRouter(1, router.NewHandler())
	server.Run()

}
