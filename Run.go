package main

import (
	"Nb/iface"
	"Nb/model"
	"Nb/net"
	"Nb/router"
	"Nb/utils"
	"log"
)

func main() {
	server := net.NewServer()
	server.AddRouter(1, router.NewHandler())
	server.AddRouter(2, router.NewAnswerHandler())
	server.AddRouter(3, router.NewAnswerHandler())
	server.SetOnConnStop(func(connection iface.IConnection) {
		id, ok := connection.GetProperty("deviceId")
		if !ok {
			return
		}
		deviceId := id.(string)
		db := utils.GlobalObject.Db
		//获取node_id
		node := model.AutoNode{}
		db.Model(&node).Where("duid=?", deviceId).Update("is_online", "N")
		log.Println(id)
	})
	server.Run()
}
