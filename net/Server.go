package net

import (
	"Nb/iface"
	"Nb/utils"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Host       string
	Port       int
	msgHandler iface.IMsgHandler
	manager    iface.IConnManager
	startHook  func(connection iface.IConnection)
	stopHook   func(connection iface.IConnection)
	protocol   iface.IProtocol
}

func (s *Server) SetProtocol(protocol iface.IProtocol) {
	s.protocol = protocol
}

func (s *Server) GetProtocol() iface.IProtocol {
	return s.protocol
}

func (s *Server) SetOnConnStart(hook func(connection iface.IConnection)) {
	s.startHook = hook
}

func (s *Server) CallOnConnStart(connection iface.IConnection) {
	if s.startHook != nil {
		s.startHook(connection)
	}
}

func (s *Server) SetOnConnStop(hook func(connection iface.IConnection)) {
	s.stopHook = hook
}

func (s *Server) CallOnConnStop(connection iface.IConnection) {
	if s.stopHook != nil {
		s.stopHook(connection)
	}
}

func NewServer() iface.IServer {
	return &Server{
		Host:       utils.GlobalObject.Host,
		Port:       utils.GlobalObject.Port,
		msgHandler: NewMsgHandler(),
		manager:    NewConnManager(),
	}
}
func (s *Server) checkPing() {
	ticker := time.NewTicker(time.Second * 10)
	for now := range ticker.C {
		fmt.Printf("当前连接数量%d\n", s.manager.Len())
		for id, con := range s.manager.GetConnections() {
			value, ok := con.GetProperty("timestamp")
			if !ok {
				con.SetProperty("timestamp", now.Unix())
				continue
			}
			timestamp := value.(int64)
			if now.Unix()-timestamp >= 60 {
				//当前连接已经3min么有数据包了
				//手动关闭
				fmt.Println("stop")
				s.manager.Remove(id)
				con.Stop()
			}
		}
	}
}
func (s *Server) Run() {
	//初始化db
	db, err := utils.NewDb()
	if err != nil {
		panic(err)
	}
	//定时器检查所有连接的上次发包时间.
	//go s.checkPing()

	utils.GlobalObject.Db = db
	utils.GlobalObject.Server = s

	//启动worker处理池
	s.msgHandler.StartWorkerPool()
	//查询连接数量
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
		s.manager.Add(connection)
		go connection.Start()
	}
}

func (s *Server) Stop() {
	for _, db := range utils.GlobalObject.Db {
		_ = db.Close()
	}
	s.manager.ClearConn()
	utils.LoggerObject.Write("服务器停止")
}
func (s *Server) AddRouter(msgId uint32, router iface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}
func (s *Server) GetMsgRouter() iface.IMsgHandler {
	return s.msgHandler
}

func (s *Server) GetManager() iface.IConnManager {
	return s.manager
}
