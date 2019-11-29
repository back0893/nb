package EXGMSG

import (
	"bytes"
	"encoding/binary"
)

/**
上传车注册信息
*/

type MsgRegister struct {
	PlatformId []byte //长11位的平台编号
	ProducerId []byte //长11位的终端厂商编号
	ModelType  []byte //长8位的终端型号
	TerminalId []byte //7位的终端编号
	Sim        []byte //长12位的sim卡号 gbk
}

func (msg *MsgRegister) Len() int {
	return 11 + 11 + 8 + 7 + 12
}

func (msg *MsgRegister) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
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
