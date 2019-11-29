package EXGMSG

import (
	"Nb/iface"
)

/**
实时上传车辆定位消息
*/

type MsgLocation struct {
	GNSSData
}

func NewMsgLocation() iface.IBody {
	return &MsgLocation{
		GNSSData: NewGNSSData(),
	}
}
