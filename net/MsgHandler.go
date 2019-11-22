package net

import (
	"Nb/iface"
	"Nb/utils"
	"fmt"
)

type MsgHandler struct {
	apis map[uint32]iface.IRouter
}

func NewMsgHandler() iface.IMsgHandler {
	return &MsgHandler{
		apis: make(map[uint32]iface.IRouter),
	}
}

func (handler *MsgHandler) DoMsgHandler(request iface.IRequest) {
	id := request.GetMsg().GetId()
	router, ok := handler.apis[id]
	if !ok {
		utils.LoggerObject.Write(fmt.Sprintf("%d对应的触发方法不存在", id))
		return
	}
	router.PerHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (handler *MsgHandler) AddRouter(msgId uint32, router iface.IRouter) {
	handler.apis[msgId] = router
}
