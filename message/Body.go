package message

import (
	"Nb/iface"
	"bytes"
	"encoding/binary"
)

/**
数据体的结构
*/
type Body struct {
	CarNum   []byte //长21位的车牌 gbk
	Color    byte
	DataType uint16 //子业务标识
	Length   uint32
	SubBody  iface.IBody //子业务员体
}

func (body *Body) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &body.CarNum); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &body.Color); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &body.DataType); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &body.Length); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &body.Length); err != nil {
		return err
	}
	if err := body.SubBody.UnmarshalUn(buffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (body *Body) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, &body.CarNum); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &body.Color); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &body.DataType); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &body.Length); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &body.Length); err != nil {
		return nil, err
	}
	if data, err := body.SubBody.Marshal(); err != nil {
		return nil, err
	} else {
		buffer.Write(data)
	}
	return buffer.Bytes(), nil
}
