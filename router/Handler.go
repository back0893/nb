package router

import (
	"Nb/iface"
	"Nb/message"
	"Nb/message/body"
	"Nb/utils"
	"fmt"
)

type Handler struct {
	Base
}

func NewHandler() iface.IRouter {
	return &Handler{}
}
func (hand *Handler) Handle(request iface.IRequest) {
	msg := request.GetMsg()
	utils.LoggerObject.Write(fmt.Sprintf("%d", msg.GetId()))
	header := message.Header{
		SN:          1,
		ID:          0x1002,
		UUId:        1,
		Version:     []byte{1, 2, 3},
		EncryptFlag: 0,
		EncryptKey:  0,
	}
	if value, ok := request.GetProperty("sn"); ok {
		header.SN = value.(uint32)
	} else {
		header.SN = 1
	}

	body_msg := &body.ConnectRsp{
		Result:     0x00,
		VerifyCode: 1,
	}
	response := &message.Message{
		Header: &header,
		Body:   body_msg,
	}
	request.GetConnection().SendBuffMsg(response)
}
