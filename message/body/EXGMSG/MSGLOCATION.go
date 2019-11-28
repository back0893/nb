package EXGMSG

import (
	"Nb/iface"
	"bytes"
	"encoding/binary"
)

/**
实时上传车辆定位消息
*/

type MsgLocation struct {
	EXGMSG
	GNSSData
}

func NewMsgLocation() iface.IBody {
	return &MsgLocation{
		EXGMSG:   NewEXGMSG(),
		GNSSData: NewGNSSData(),
	}
}

func (msg *MsgLocation) UnmarshalUn(data []byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &msg.Encrypt); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Date); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Time); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Lng); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Lat); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Vec1); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Vec2); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Vec3); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Direction); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Altitude); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.State); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.Alarm); err != nil {
		return err
	}
	return nil
}

func (msg *MsgLocation) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, &msg.CarNum); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Color); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.DataType); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Length); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Encrypt); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Date); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Time); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Lng); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Lat); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Vec1); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Vec2); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Vec3); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Direction); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Altitude); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.State); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.Alarm); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
