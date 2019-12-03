package router

import (
	"Nb/iface"
	"Nb/message"
)

type Jt1003 struct {
	Base
}

func (Jt1003) Handle(request iface.IRequest) {
	msg := request.GetMsg().(*message.Message)
	//生成1004的回应
	header := message.NewHeader()
	header.ID = 0x1004
	header.Version = msg.Header.Version

	if value, ok := request.GetProperty("sn"); ok {
		header.SN = value.(uint32)
	} else {
		header.SN = 1
	}

	//1004的body回应
	body_msg := message.EmptyBody{}
	response := &message.Message{
		Header: header,
		Body:   body_msg,
	}
	request.GetConnection().SendBuffMsg(response)
}
