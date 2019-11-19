package net

import (
	"Nb/iface"
	"Nb/utils"
	"fmt"
	"net"
)

type Server struct {
	Host string
	Port int
}

func NewServer() iface.IServer {
	return &Server{
		Host: "0.0.0.0",
		Port: 8001,
	}
}
func (s *Server) Run() {
	s.Server()
}

func (s *Server) Server() {
	host := fmt.Sprintf("%s:%d", s.Host, s.Port)
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return
	}
	server, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return
	}
	for {
		conn, err := server.AcceptTCP()
		if err != nil {
			utils.LoggerObject.Write(err.Error())
			return
		}
		go func(conn *net.TCPConn) {
			data := make([]byte, 1024)
			n, err := conn.Read(data)
			if err != nil {
				utils.LoggerObject.Write(err.Error())
			}
			utils.LoggerObject.Write(string(data[:n]))
			conn.Write([]byte("ok"))
		}(conn)
	}
}

func (Server) Stop() {
	utils.LoggerObject.Write("服务器停止")
}
