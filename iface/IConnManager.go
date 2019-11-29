package iface

type IConnManager interface {
	Add(connection IConnection)
	GetConnection(uint64) (IConnection, bool)
	Remove(uint64)
	Len() int
	ClearConn()
	GetConnections() map[uint64]IConnection
}
