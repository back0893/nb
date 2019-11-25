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
	utils.LoggerObject.Write(fmt.Sprintf("%d", request.GetMsg().GetId()))
	request.GetConnection().SendBuffMsg([]byte{'o', 'k'})
}
