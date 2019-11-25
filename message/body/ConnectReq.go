package body

import (
	"bytes"
	"encoding/binary"
)

type ConnectReq struct {
	UserId       uint32
	Password     []byte //创建时设置长度
	DownLinkIp   []byte //创建时设置长度
	DownLinkPort uint16
}

func (req *ConnectReq) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &req.UserId); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &req.Password); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &req.DownLinkIp); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &req.DownLinkPort); err != nil {
		return err
	}
	return nil
}

func (req *ConnectReq) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, req.UserId); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, req.Password); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, req.DownLinkIp); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, req.DownLinkPort); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
