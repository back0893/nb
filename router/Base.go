package router

import "Nb/iface"

type Base struct{}

func (Base) PerHandle(request iface.IRequest) {}

func (Base) Handle(request iface.IRequest) {}

func (Base) PostHandle(request iface.IRequest) {}
