package router

import (
	"Nb/iface"
	"Nb/utils"
)

type Handler struct {
	Base
}

func NewHandler() iface.IRouter {
	return &Handler{}
}
func (hand *Handler) Handle(request iface.IRequest) {
	utils.LoggerObject.Write(string(request.GetData()))
	request.GetConnection().Write(request.GetData())
}
