package net

import "Nb/iface"

type Request struct {
	connection iface.IConnection
	data       []byte
}

func NewRequest(connection iface.IConnection, data []byte) iface.IRequest {
	return &Request{
		connection: connection,
		data:       data,
	}
}

func (r *Request) GetConnection() iface.IConnection {
	return r.connection
}

func (r *Request) GetData() []byte {
	return r.data
}
