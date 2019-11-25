package iface

import "net"

type IConnection interface {
	GetConn() *net.TCPConn
	Write([]byte) (int, error) //直接使用连接发送
	GetConId() uint64
	Start()
	Stop()
	SendMsg([]byte)                             //读写分离
	SendBuffMsg([]byte)                         //读写分离
	SetProperty(key string, value interface{})  //设置属性
	GetProperty(key string) (interface{}, bool) //读取
	RemoveProperty(key string)
}
