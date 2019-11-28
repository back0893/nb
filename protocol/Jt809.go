package protocol

import (
	"Nb/iface"
	"Nb/message"
	"Nb/message/body/EXGMSG"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type Jt809 struct{}

func (Jt809) SplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 && atEOF == true {
		return 0, nil, io.EOF
	}
	start_index := bytes.IndexByte(data, 0x5b)
	if start_index == -1 {
		//灭有开始标志,跳过
		return len(data), nil, nil
	}
	end_index := bytes.IndexByte(data, 0x5d)
	if end_index == -1 && atEOF == true {
		//数据不全导致
		return 0, nil, io.EOF
	}
	if start_index >= end_index {
		//上一个包异常
		return start_index, nil, nil
	}
	message := data[start_index+1 : end_index]
	return end_index + 1, message, nil
}
func (protocol *Jt809) getDataType(data *[]byte) (uint16, error) {
	buffer := bytes.NewBuffer(*data)
	buffer.Next(22 + 22)
	var data_type uint16
	if err := binary.Read(buffer, binary.BigEndian, &data_type); err != nil {
		return 0, err
	}
	return data_type, nil
}
func (protocol *Jt809) getID(data *[]byte) (uint16, error) {
	buffer := bytes.NewBuffer(*data)
	buffer.Next(8)
	var id uint16
	if err := binary.Read(buffer, binary.BigEndian, &id); err != nil {
		return 0, err
	}
	return id, nil
}

func (protocol *Jt809) Decode(data []byte) (iface.IMessage, error) {
	//转义
	/**
	在数据发送时进行了如下的转义
	转义 0x5b=>0x5a 0x01
	转义 0x5a=>0x5a 0x02
	转义 0x5d=>0x5e 0x01
	转义 0x5e=>0x5e 0x01
	*/
	data = bytes.ReplaceAll(data, []byte{0x5a, 0x01}, []byte{0x5b})
	data = bytes.ReplaceAll(data, []byte{0x5a, 0x02}, []byte{0x5a})
	data = bytes.ReplaceAll(data, []byte{0x5e, 0x01}, []byte{0x5d})
	data = bytes.ReplaceAll(data, []byte{0x5e, 0x02}, []byte{0x5e})
	id, err := protocol.getID(&data)
	if err != nil {
		return nil, err
	}
	switch id {
	case 0x1001:
		//使用messgae
		msg := message.NewMessage()
		if err := msg.UnmarshalUn(data); err != nil {
			return nil, err
		}
		return msg, nil
	case 0x1200:
		//这里需要进一步判断data_type
		data_type, err := protocol.getDataType(&data)
		if err != nil {
			return nil, err
		}
		switch data_type {
		case 0x1201:
		case 0x1202:
			msg := message.Message{
				Header: message.NewHeader(),
				Body:   EXGMSG.NewMsgLocation(),
			}
			if err := msg.UnmarshalUn(data); err != nil {
				return nil, err
			}
			//重新设置消息id
			return &msg, err
		case 0x1203:

		}
	}
	return nil, errors.New("灭有寻找到")
}

func (Jt809) Encode(msg iface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{0x5b})
	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}
	data = bytes.ReplaceAll(data, []byte{0x5b}, []byte{0x5a, 0x01})
	data = bytes.ReplaceAll(data, []byte{0x5a}, []byte{0x5a, 0x02})
	data = bytes.ReplaceAll(data, []byte{0x5d}, []byte{0x5e, 0x01})
	data = bytes.ReplaceAll(data, []byte{0x5e}, []byte{0x5e, 0x02})
	buffer.Write(data)
	buffer.WriteByte(0x5d)
	return buffer.Bytes(), nil
}
