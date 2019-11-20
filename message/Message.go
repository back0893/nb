package message

import (
	"bytes"
	"strconv"
	"strings"
)

type Message struct {
	Version  string // 软件版本
	DeviceId string //设备id
	Switch1  string //开端输入1
	Switch2  string //开端输入1
	Input1   string //模拟输入1
	Input2   string //模拟输入2
	Voltage  string //电压
	Record   int    //采样频率
	Upload   int    //上传频率
	Ack      int    //16进制的校验码
	rawData  []byte //原始数据
}

func NewMessage() *Message {
	return &Message{}
}

func (msg *Message) GetData() []byte {
	return msg.rawData
}
func (msg *Message) UnmarshalUn(data []byte) error {
	splitData := strings.Split(string(data), ":")

	msg.Version = splitData[0]

	msg.DeviceId = splitData[1]

	msg.Switch1 = splitData[2]

	msg.Switch2 = splitData[3]

	msg.Input1 = splitData[4]

	msg.Input2 = splitData[5]

	msg.Voltage = splitData[6]

	if s, err := strconv.Atoi(splitData[7]); err != nil {
		return err
	} else {
		msg.Record = s
	}

	if s, err := strconv.Atoi(splitData[8]); err != nil {
		return err
	} else {
		msg.Upload = s
	}

	if s, err := strconv.Atoi(splitData[9]); err != nil {
		return err
	} else {
		msg.Ack = s
	}
	msg.rawData = data
	return nil
}

func (msg *Message) CheckAck() bool {
	needCheck := msg.rawData[:len(msg.rawData)-1]
	ack := 0
	for _, ord := range needCheck {
		ack += int(ord)
	}
	return ack == msg.Ack
}

func (msg *Message) Marshal() []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte('{')
	data := []string{"A"}
	data = append(data, msg.Switch1)
	data = append(data, msg.Switch2)
	data = append(data, msg.Input1)
	data = append(data, msg.Input2)
	data = append(data, msg.Voltage)
	buffer.WriteString(strings.Join(data, ":"))
	buffer.WriteByte('}')
	return buffer.Bytes()
}
