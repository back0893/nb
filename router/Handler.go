package router

import (
	"Nb/iface"
	"log"
)

type Handler struct {
	Base
}

func NewHandler() iface.IRouter {
	return &Handler{}
}
func (hand *Handler) Handle(request iface.IRequest) {
	log.Println(request.GetMsg())
	request.GetConnection().Write(request.GetMsg().Marshal())
}
