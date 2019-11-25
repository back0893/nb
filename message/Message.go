package message

import (
	"Nb/iface"
	"Nb/message/body"
	"bytes"
	"encoding/binary"
)

type Header struct {
	Len         uint32
	SN          uint32
	ID          uint16
	UUId        uint32
	Version     [3]byte
	EncryptFlag byte
	EncryptKey  uint32
}

func (header *Header) UnmarshalUn([]byte) error {
	panic("implement me")
}

func (header *Header) Marshal() ([]byte, error) {
	panic("implement me")
}

type Message struct {
	Header  Header
	Body    iface.IBody
	Crc     uint16
	rawData []byte
}

func NewMessage() iface.IMessage {
	return &Message{
		Header: Header{},
		Body:   &body.ConnectReq{},
	}
}
func (msg *Message) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := msg.Header.UnmarshalUn(buffer.Next(22)); err != nil {
		return err
	}
	body_len := msg.Header.Len - 22 - 1
	if err := msg.Body.UnmarshalUn(buffer.Next(int(body_len))); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Crc); err != nil {
		return err
	}
	return nil
}

func (msg *Message) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	header_data, err := msg.Header.Marshal()
	if err != nil {
		return nil, err
	}
	buffer.Write(header_data)
	body_data, err := msg.Body.Marshal()
	if err != nil {
		return nil, err
	}
	buffer.Write(body_data)
	if err := binary.Write(buffer, binary.BigEndian, msg.Crc); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (msg *Message) GetData() []byte {
	return msg.rawData
}

func (msg *Message) GetId() uint32 {
	return uint32(msg.Header.ID)
}

func (Message) String() string {
	return "message"
}
