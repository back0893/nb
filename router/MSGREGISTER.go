package router

import "Nb/iface"

type MSGREGISTER struct {
	Base
}

func (MSGREGISTER) Handle(request iface.IRequest) {
}
