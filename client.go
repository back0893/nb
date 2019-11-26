package main

import (
	"Nb/message"
	"Nb/message/body"
	"Nb/protocol"
	"log"
	"net"
)

func main() {
	header := message.Header{
		Len:         0,
		SN:          1,
		ID:          0x1001,
		UUId:        1,
		Version:     []byte{1, 2, 3},
		EncryptFlag: 0,
		EncryptKey:  0,
	}
	ip := make([]byte, 32, 32)
	copy(ip, []byte("192.168.1.1"))
	body_msg := &body.ConnectReq{
		UserId:       1,
		Password:     []byte{1, 2, 3, 4, 5, 6, 7, 8},
		DownLinkIp:   ip,
		DownLinkPort: 3005,
	}
	msg := message.Message{
		Header: header,
		Body:   body_msg,
		Crc:    1,
	}
	data, err := msg.Marshal()
	if err != nil {
		panic(err)
	}
	jt809 := protocol.Jt809{}
	log.Printf("%x", jt809.Encode(data))
	return
	//todo 发送
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10034")
	if err != nil {
		panic(err)
	}
	client, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	client.Write(jt809.Encode(data))
}
