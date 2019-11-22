package iface

import "net"

type IConnection interface {
	GetConn() *net.TCPConn
	Write([]byte) (int, error) //直接使用连接发送
	GetConId() uint64
	Start()
	Stop()
	SendMsg([]byte)     //读写分离
	SendBuffMsg([]byte) //读写分离
}
