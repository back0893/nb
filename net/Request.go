package net

import "Nb/iface"

type Request struct {
	connection iface.IConnection
	msg        iface.IMessage
}

func NewRequest(connection iface.IConnection, message iface.IMessage) iface.IRequest {
	return &Request{
		connection: connection,
		msg:        message,
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
