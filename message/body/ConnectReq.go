package body

/**
连接请求
*/
import (
	"Nb/iface"
	"Nb/utils"
	"bytes"
	"encoding/binary"
)

type ConnectReq struct {
	UserId       uint32
	Password     []byte //创建时设置长度8 gbk
	DownLinkIp   []byte //创建时设置长度32 gbk
	DownLinkPort uint16
}

func (req *ConnectReq) Len() int {
	return 4 + 8 + 32 + 2
}

func NewConnectReq() iface.IBody {
	return &ConnectReq{
		Password:   make([]byte, 8, 8),
		DownLinkIp: make([]byte, 32, 32),
	}
}

func (req *ConnectReq) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &req.UserId); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &req.Password); err != nil {
		return err
	}
	req.Password = utils.GlobalObject.ConvertToString(req.Password, "gbk", "utf-8")
	if err := binary.Read(buffer, binary.BigEndian, &req.DownLinkIp); err != nil {
		return err
	}
	req.DownLinkIp = utils.GlobalObject.ConvertToString(req.DownLinkIp, "gbk", "utf-8")
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
	req.Password = utils.GlobalObject.ConvertToString(req.Password, "utf-8", "gbk")
	if err := binary.Write(buffer, binary.BigEndian, req.Password); err != nil {
		return nil, err
	}
	req.DownLinkIp = utils.GlobalObject.ConvertToString(req.DownLinkIp, "utf-8", "gbk")
	if err := binary.Write(buffer, binary.BigEndian, req.DownLinkIp); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, req.DownLinkPort); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}