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

	header_msg := message.Header{
		Len:         0,
		SN:          1,
		ID:          0x1001,
		UUId:        1,
		Version:     []byte{1, 2, 3},
		EncryptFlag: 0,
		EncryptKey:  0,
	}
	body_msg := body.ConnectRsp{
		Result:     0x00,
		VerifyCode: 1,
	}
	response := message.Message{
		Header: header_msg,
		Body:   body_msg,
	}
	data, err := response.Marshal()
	if err != nil {
		utils.LoggerObject.Write(err.Error())
		return
	}
	request.GetConnection().SendBuffMsg(data)
}
