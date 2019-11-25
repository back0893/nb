package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func main() {
	str := "hiworld"
	buffer := bytes.NewBufferString(str)
	r1 := make([]byte, 2)
	r2 := make([]byte, 5)
	binary.Read(buffer, binary.BigEndian, &r1)
	binary.Read(buffer, binary.BigEndian, &r2)
	log.Println(string(r1))
	log.Println(string(r2))
	//header:=message.Header{
	//	Len:0,
	//	SN:1,
	//	ID:1,
	//	UUId:1,
	//	Version:[3]byte{0x01,0x01,0x01},
	//	EncryptFlag:0,
	//	EncryptKey:0,
	//}
	//body_msg:=&body.ConnectReq{
	//	UserId:1,
	//	Password:[]byte{1,2,3,4,5,6,7,8},
	//	DownLinkIp:[]byte{'1','9','2'},
	//	DownLinkPort:3005,
	//}
	//msg:=message.Message{
	//	Header:header,
	//	Body:body_msg,
	//	Crc:1,
	//}
	//data,err:=msg.Marshal()
	//if err!=nil{
	//	panic(err)
	//}
	//todo 发送

}
