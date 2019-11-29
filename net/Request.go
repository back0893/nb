package net

import (
	"Nb/iface"
	"sync"
)

type Request struct {
	connection iface.IConnection
	msg        iface.IMessage
	//设置一些临时的请求数据
	lock     sync.RWMutex
	property map[string]interface{}
}

func NewRequest(connection iface.IConnection, message iface.IMessage) iface.IRequest {
	return &Request{
		connection: connection,
		msg:        message,
		property:   make(map[string]interface{}),
	}
}

func (r *Request) GetConnection() iface.IConnection {
	return r.connection
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsg() iface.IMessage {
	return r.msg
}
func (c *Request) SetProperty(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.property[key] = value

}

func (c *Request) GetProperty(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	t, ok := c.property[key]
	return t, ok
}
