package iface

type IRouter interface {
	PerHandle(request IRequest)
	Handle(request IRequest)
	PostHandle(request IRequest)
}
