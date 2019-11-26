package router

import (
	"Nb/iface"
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
}
