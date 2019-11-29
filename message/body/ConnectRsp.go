package body

/**
连接回应
*/
import (
	"bytes"
	"encoding/binary"
)

type ConnectRsp struct {
	Result     byte
	VerifyCode uint32
}

func (msg *ConnectRsp) Len() int {
	return 5
}

func NewConnectRsp() *ConnectRsp {
	return &ConnectRsp{}
}

func (msg *ConnectRsp) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &msg.Result); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.VerifyCode); err != nil {
		return err
	}
	return nil
}

func (msg *ConnectRsp) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, msg.Result); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, msg.VerifyCode); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
