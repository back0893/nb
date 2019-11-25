package router

import (
	"Nb/iface"
	"Nb/utils"
)

type AnswerHandler struct {
	Base
}

func NewAnswerHandler() iface.IRouter {
	return &AnswerHandler{}
}

func (AnswerHandler) Handle(request iface.IRequest) {
	utils.LoggerObject.Write(request.GetMsg().String())
	request.GetConnection().SendBuffMsg(request.GetMsg().Marshal())
}
