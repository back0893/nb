package body

import (
	"bytes"
	"encoding/binary"
)

/**
上传车注册信息
*/

type MsgRegister struct {
	CarNum     []byte //长21位的车牌 gbk
	Color      byte
	DataType   uint16 //子业务标识
	Length     uint32
	PlatformId []byte //长11位的平台编号
	ProducerId []byte //长11位的终端厂商编号
	ModelType  []byte //长8位的终端型号
	TerminalId []byte //7位的终端编号
	Sim        []byte //长12位的sim卡号 gbk
}

func (msg *MsgRegister) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &msg.CarNum); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Color); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.DataType); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Length); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.PlatformId); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.ProducerId); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.TerminalId); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Sim); err != nil {
		return err
	}
	return nil
}

func (msg MsgRegister) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, &msg.CarNum); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Color); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.DataType); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Length); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.PlatformId); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.ProducerId); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.TerminalId); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Sim); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
