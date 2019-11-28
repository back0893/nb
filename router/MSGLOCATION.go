package router

import (
	"Nb/iface"
	"Nb/message"
	"Nb/message/body/EXGMSG"
	"Nb/model"
	"Nb/utils"
	"fmt"
	"log"
	"time"
)

type EXGMSGRouter struct {
	Base
}

func NewEXGMSGRouter() iface.IRouter {
	return &EXGMSGRouter{}
}

func (router *EXGMSGRouter) Handle(request iface.IRequest) {
	date := time.Now()
	exg_body := request.GetMsg().(*message.Message).Body
	switch exg_body.(type) {
	case *EXGMSG.MsgLocation:
		location := request.GetMsg().(*message.Message).Body.(*EXGMSG.MsgLocation)
		log.Println(location)
		point := model.GpsPoint{
			PlateNum:  string(location.CarNum),
			PointX:    float32(location.Lat) / 10e6,
			PointY:    float32(location.Lng) / 10e6,
			VehicleId: 0,
			Height:    uint(location.Altitude),
			Speed:     float32(location.Vec2),
			Direction: uint(location.Direction),
			SimCard:   "", //目前没有
			Time:      uint(date.Unix()),
		}
		utils.LoggerObject.Write(fmt.Sprintf("lat=>%.6f,lng=>%.6f", point.PointX, point.PointY))
		db, ok := utils.GlobalObject.Db["gps"]
		if !ok {
			utils.LoggerObject.Write("!gps连接失败!")
			return
		}
		table_name := fmt.Sprintf("vehicle_point_%s", date.Format("2006-01-02"))
		db.Table(table_name).Create(&point)
	case *EXGMSG.MsgRegister:
	case *EXGMSG.HISTORYARCOSSAREA:
	}
	utils.LoggerObject.Write("!location!")
}
