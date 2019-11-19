package iface

import "net"

type IConnection interface {
	GetConn() *net.TCPConn
	Write([]byte) (int, error)
	GetConId() uint64
	Start()
	Stop()
}
