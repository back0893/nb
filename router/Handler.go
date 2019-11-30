package router

import (
	"Nb/iface"
	"Nb/message"
	"Nb/model"
	"Nb/utils"
	"fmt"
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
	iMessage := request.GetMsg()
	msg := iMessage.(*message.Message)
	//设置连接的属性
	request.GetConnection().SetProperty("deviceId", msg.DeviceId)
	hand.doMsg(msg)
	request.GetConnection().SendBuffMsg(request.GetMsg().Marshal())
}
func (hand *Handler) doMsg(msg *message.Message) {
	db := utils.GlobalObject.Db
	//获取node_id
	node := model.AutoNode{}
	db.Unscoped().Where("duid=?", msg.DeviceId).First(&node)
	if node.ID == 0 {
		//没有对应的id不做任何操作
		utils.LoggerObject.Write(fmt.Sprintf("%s设备没有配置", msg.DeviceId))
		return
	}
	//更新设备在线
	//node.IsOnline = "Y"
	//更新
	//db.Model(&node).Where("duid=?", msg.DeviceId).Update("is_online", "Y")
	//有node_id 更新pin表中的实时数据
	pins := make([]model.AutoNodePing, 0)
	db.Unscoped().Where("node_id=?", node.ID).Find(&pins)
	if len(pins) == 0 {
		//没有就不做更新操作
		utils.LoggerObject.Write(fmt.Sprintf("%s设备没有配置pin表", msg.DeviceId))
		return
	}
	for _, pin := range pins {
		utils.LoggerObject.Write(pin.Time.Format("2006-01-02T15:04:05"))
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
			val, err := strconv.Atoi(msg.Input1[1:])
			if err != nil {
				continue
			}
			pin.PinValue = val
			pin.Unit = "mV"
		} else if pin.PinId == 4 {
			//开关输入4
			val, err := strconv.Atoi(msg.Input2[1:])
			if err != nil {
				continue
			}
			pin.PinValue = val
			pin.Unit = "mV"
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
		//对于开关量有一个额外的alarm参数
		if pin.Type == 0 {
			node_value.Alarm = node_value.Value
		}
		db.Create(&node_value)
	}
}
