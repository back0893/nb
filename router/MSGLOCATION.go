package router

import (
	"Nb/iface"
	"Nb/message"
	"Nb/message/body"
	"Nb/model"
	"Nb/utils"
	"fmt"
	"time"
)

type MSGLOCATION struct {
	Base
}

func (router *MSGLOCATION) Handle(request iface.IRequest) {
	date := time.Now()
	location := request.GetMsg().(*message.Message).Body.(*body.MsgLocation)
	point := model.GpsPoint{
		PlateNum:  string(location.CarNum),
		PointX:    float32(location.Lat / 10e6),
		PointY:    float32(location.Lng / 10e6),
		VehicleId: 0,
		Height:    uint(location.Altitude),
		Speed:     float32(location.Vec2),
		Direction: uint(location.Direction),
		SimCard:   "", //目前没有
		Time:      uint(date.Unix()),
	}
	db, ok := utils.GlobalObject.Db["gps"]
	if !ok {
		utils.LoggerObject.Write("!gps连接失败!")
		return
	}
	table_name := fmt.Sprintf("vehicle_point_%s", date.Format("2006-01-02"))
	db.Table(table_name).Create(&point)
}
