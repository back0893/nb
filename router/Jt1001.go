package router

import (
	"Nb/iface"
	"Nb/message"
	"Nb/message/body"
	"Nb/utils"
	"fmt"
)

type Jt1001 struct {
	Base
}

func NewJt1001() iface.IRouter {
	return &Jt1001{}
}
func (hand *Jt1001) Handle(request iface.IRequest) {
	msg := request.GetMsg().(*message.Message)
	utils.LoggerObject.Write(fmt.Sprintf("%d", msg.GetId()))
	header := message.MakeHeader(0x1002, msg.Header.Version)
	if value, ok := request.GetConnection().GetProperty("sn"); ok {
		header.SN = value.(uint32)
	} else {
		header.SN = 1
	}

	//1002的body回应
	body_msg := body.NewConnectRsp()
	body_msg.VerifyCode = 1
	body_msg.Result = 0x00

	response := &message.Message{
		Header: header,
		Body:   body_msg,
	}
	request.GetConnection().SetProperty("login", true)
	request.GetConnection().SendBuffMsg(response)

}
