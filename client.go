package main

import (
	"Nb/message"
	"Nb/message/body"
	"Nb/message/body/EXGMSG"
	"Nb/protocol"
	"log"
	"net"
)

func jt1001() []byte {
	header := message.Header{
		Len:         0,
		SN:          1,
		ID:          0x1001,
		UUId:        1,
		Version:     []byte{1, 2, 3},
		EncryptFlag: 0,
		EncryptKey:  1,
	}
	password := make([]byte, 8, 8)
	a := []byte("中文")
	//a:=utils.GlobalObject.ConvertToString([]byte("123456"),"utf-8","gbk")
	copy(password, a)
	ip := make([]byte, 32, 32)
	copy(ip, []byte("192.168.1.1"))
	body_msg := &body.ConnectReq{
		UserId:       1,
		Password:     password,
		DownLinkIp:   ip,
		DownLinkPort: 3005,
	}
	msg := &message.Message{
		Header: &header,
		Body:   body_msg,
	}

	jt809 := protocol.Jt809{}
	raw, err := jt809.Encode(msg)
	if err != nil {
		panic(err)
	}
	return raw

}
func jt1202() []byte {
	header := message.Header{
		Len:         0,
		SN:          1,
		ID:          0x1200,
		UUId:        1,
		Version:     []byte{1, 2, 3},
		EncryptFlag: 0,
		EncryptKey:  1,
	}
	car_num := make([]byte, 21)
	copy(car_num, []byte("川a123456"))
	gnss := EXGMSG.GNSSData{
		Encrypt:   0,
		Date:      []byte{21, 11, 0x07, 0xe3},
		Time:      []byte{19, 06, 02},
		Lng:       103000000,
		Lat:       301111111,
		Vec1:      1,
		Vec2:      2,
		Vec3:      100,
		Direction: 150,
		Altitude:  120,
		State:     0,
		Alarm:     0,
	}
	body_msg := &message.Body{
		CarNum:   car_num,
		Color:    1,
		DataType: 0x1202,
		Length:   0,
		SubBody: &EXGMSG.MsgLocation{
			GNSSData: gnss,
		},
	}

	msg := &message.Message{
		Header: &header,
		Body:   body_msg,
	}
	jt809 := protocol.Jt809{}
	raw, err := jt809.Encode(msg)
	if err != nil {
		panic(err)
	}
	return raw

}
func main() {

	//log.Printf("%x", raw)
	//return
	//todo 发送
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10034")
	if err != nil {
		panic(err)
	}
	client, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	jt1001 := jt1001()
	client.Write(jt1001)
	rev := make([]byte, 1024)
	i, err := client.Read(rev)
	if err != nil {
		panic(err)
	}
	log.Printf("%x", rev[:i])
	jt1202 := jt1202()
	client.Write(jt1202)
	rev = make([]byte, 1024)
	i, err = client.Read(rev)
	if err != nil {
		panic(err)
	}
	log.Printf("%x", rev[:i])
}
