package router

import (
	"Nb/iface"
	"Nb/message"
	"Nb/model"
	"Nb/utils"
	"strconv"
	"time"
)

type Handler struct {
	Base
}

func NewHandler() iface.IRouter {
	return &Handler{}
}
func (hand *Handler) Handle(request iface.IRequest) {
	hand.doMsg(request.GetMsg())
	request.GetConnection().Write(request.GetMsg().Marshal())
}
func (hand *Handler) doMsg(iMessage iface.IMessage) {
	msg := iMessage.(*message.Message)
	db := utils.GlobalObject.Db
	//获取node_id
	node := model.AutoNode{}
	db.Unscoped().Where("duid=?", msg.DeviceId).First(&node)
	if node.ID == 0 {
		//没有对应的id不做任何操作
		return
	}
	//有node_id 更新pin表中的实时数据
	pins := make([]model.AutoNodePing, 0)
	db.Unscoped().Where("node_id=?", node.ID).Find(&pins)
	if len(pins) == 0 {
		//没有就不做更新操作
		return
	}
	for _, pin := range pins {
		if pin.PinId == 1 {
			//开关输入1
			val, err := strconv.Atoi(msg.Switch1[1:])
			if err != nil {
				continue
			}
			pin.PinValue = val
		} else if pin.PinId == 2 {
			//开关输入2
			val, err := strconv.Atoi(msg.Switch2[1:])
			if err != nil {
				continue
			}
			pin.PinValue = val
		} else if pin.PinId == 3 {
			//开关输入3
			val, err := strconv.ParseFloat(msg.Input1[1:6], 32)
			if err != nil {
				continue
			}
			pin.PinValue = int(val * 1000)
			pin.Unit = "V"
		} else if pin.PinId == 4 {
			//开关输入4
			val, err := strconv.ParseFloat(msg.Input2[1:6], 32)
			if err != nil {
				continue
			}
			pin.PinValue = int(val * 1000)
			pin.Unit = "V"
		}
		pin.Time = time.Now()
		db.Save(&pin)
		//再去插入value中的历史数据
		node_value := model.AutoNodeValue{}
		node_value.Time = pin.Time
		node_value.NodeId = int(node.ID)
		node_value.Type = pin.Type
		node_value.Value = pin.PinValue
		node_value.Pin = int(pin.ID)
		db.Create(&node_value)
	}
}
