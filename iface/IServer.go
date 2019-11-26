package iface

type IServer interface {
	Run()
	Server()
	Stop()
	AddRouter(msgId uint32, router IRouter)
	GetMsgRouter() IMsgHandler
	GetManager() IConnManager
	SetOnConnStart(func(connection IConnection))
	CallOnConnStart(connection IConnection)
	SetOnConnStop(func(connection IConnection))
	CallOnConnStop(connection IConnection)
	SetProtocol(protocol IProtocol)
	GetProtocol() IProtocol
}
