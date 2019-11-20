package net

import (
	"strconv"
	"strings"
)

type Message struct {
	Version  string  // 软件版本
	DeviceId string  //设备id
	Switch1  int     //开端输入1
	Switch2  int     //开端输入1
	Input1   float32 //模拟输入1
	Input2   float32 //模拟输入2
	Voltage  float32 //电压
	Ack      int     //16进制的校验码
	rawData  []byte  //原始数据
}

func (msg *Message) UnmarshalUn(data []byte) error {
	splitData := strings.Split(string(data), ":")

	msg.Version = splitData[0]

	msg.DeviceId = splitData[1]

	if s, err := strconv.Atoi(splitData[2]); err != nil {
		return err
	} else {
		msg.Switch1 = s
	}

	if s, err := strconv.Atoi(splitData[3]); err != nil {
		return err
	} else {
		msg.Switch2 = s
	}

	if s, err := strconv.ParseFloat(splitData[4], 32); err != nil {
		return err
	} else {
		msg.Input1 = float32(s)
	}

	if s, err := strconv.ParseFloat(splitData[5], 32); err != nil {
		return err
	} else {
		msg.Input2 = float32(s)
	}
	if s, err := strconv.ParseFloat(splitData[6], 32); err != nil {
		return err
	} else {
		msg.Voltage = float32(s)
	}
	if s, err := strconv.Atoi(splitData[7]); err != nil {
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
