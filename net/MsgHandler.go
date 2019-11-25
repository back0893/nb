package net

import (
	"Nb/iface"
	"Nb/utils"
	"fmt"
)

type MsgHandler struct {
	apis               map[uint32]iface.IRouter
	MaxWorkerSize      int
	MaxWorkerQueueTask int
	workerPool         []chan iface.IRequest
}

func (handler *MsgHandler) StartOneWorker(taskQueue chan iface.IRequest) {
	for {
		select {
		case request := <-taskQueue:
			handler.DoMsgHandler(request)
		}
	}
}
func (handler *MsgHandler) StartWorkerPool() {
	for i := 0; i < handler.MaxWorkerSize; i++ {
		handler.workerPool[i] = make(chan iface.IRequest, handler.MaxWorkerQueueTask)
		go handler.StartOneWorker(handler.workerPool[i])
	}
}

func (handler *MsgHandler) SendMsgToTaskQueue(request iface.IRequest) {
	conId := request.GetMsg().GetId()
	index := int(conId) % handler.MaxWorkerSize
	handler.workerPool[index] <- request

}

func NewMsgHandler() iface.IMsgHandler {
	return &MsgHandler{
		apis:               make(map[uint32]iface.IRouter),
		workerPool:         make([]chan iface.IRequest, utils.GlobalObject.MaxWorkerSize),
		MaxWorkerSize:      utils.GlobalObject.MaxWorkerSize,
		MaxWorkerQueueTask: utils.GlobalObject.MaxWorkerQueueTask,
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
