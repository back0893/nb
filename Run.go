package main

import (
	"Nb/net"
	"Nb/router"
)

func main() {
	server := net.NewServer()
	server.AddRouter(1, router.NewHandler())
	server.AddRouter(2, router.NewAnswerHandler())
	server.AddRouter(3, router.NewAnswerHandler())
	server.Run()
}
