package router

import (
	"Nb/iface"
	"Nb/message"
)

/**
主链路的连接保存请求
*/

type Jt1005 struct {
	Base
}

func NewJt1005() iface.IRouter {
	return &Jt1005{}
}
func (Jt1005) Handle(request iface.IRequest) {
	//心跳回复
	header := request.GetMsg().(*message.Message).Header
	header.ID = 0x1006
	if value, ok := request.GetProperty("sn"); ok {
		header.SN = value.(uint32)
	} else {
		header.SN = 1
	}
	msg := message.Message{
		Header: header,
		Body:   message.NewEmptyBody(),
	}
	request.GetConnection().SendMsg(&msg)
}
