package net

import (
	"Nb/iface"
)

type ConnManager struct {
	connections map[uint64]iface.IConnection
}

func (manager *ConnManager) GetConnections() map[uint64]iface.IConnection {
	return manager.connections
}

func NewConnManager() iface.IConnManager {
	return &ConnManager{
		connections: make(map[uint64]iface.IConnection),
	}
}
func (manager *ConnManager) Add(connection iface.IConnection) {
	manager.connections[connection.GetConId()] = connection
}

func (manager *ConnManager) GetConnection(id uint64) (iface.IConnection, bool) {
	conn, ok := manager.connections[id]
	return conn, ok
}

func (manager *ConnManager) Remove(id uint64) {
	delete(manager.connections, id)
}

func (manager *ConnManager) Len() int {
	return len(manager.connections)
}

func (manager *ConnManager) ClearConn() {
	for id, conn := range manager.connections {
		conn.Stop()
		delete(manager.connections, id)
	}
}
