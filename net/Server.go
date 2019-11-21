package net

import (
	"Nb/iface"
	"Nb/utils"
	"fmt"
	"net"
)

type Server struct {
	Host   string
	Port   int
	router iface.IRouter
}

func NewServer() iface.IServer {
	return &Server{
		Host: utils.GlobalObject.Host,
		Port: utils.GlobalObject.Port,
	}
}
func (s *Server) Run() {
	//初始化db
	db, err := utils.NewDb()
	if err != nil {
		panic("!连接数据库失败!")
	}

	utils.GlobalObject.Db = db
	utils.GlobalObject.Server = s

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
	utils.LoggerObject.Write(fmt.Sprintf("%s启动成功", host))
	var connId uint64 = 0
	for {
		conn, err := server.AcceptTCP()
		connId++
		if err != nil {
			utils.LoggerObject.Write(err.Error())
			return
		}
		connection := NewConnection(connId, conn, s)
		go connection.Start()
	}
}

func (s *Server) Stop() {
	utils.GlobalObject.Db.Close()
	utils.LoggerObject.Write("服务器停止")
}
func (s *Server) AddRouter(router iface.IRouter) {
	s.router = router
}
func (s *Server) GetRouter() iface.IRouter {
	return s.router
}
