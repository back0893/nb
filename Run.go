package main

import (
	"Nb/net"
	"Nb/router"
)

func main() {
	server := net.NewServer()
	server.AddRouter(router.NewHandler())
	server.Run()

}
