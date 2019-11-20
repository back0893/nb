package iface

type IServer interface {
	Run()
	Server()
	Stop()
	AddRouter(router IRouter)
	GetRouter() IRouter
}
