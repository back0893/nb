package EXGMSG

import (
	"bytes"
	"encoding/binary"
)

type GNSSData struct {
	Encrypt   byte   //是否加密
	Date      []byte //4位长的 dmyy
	Time      []byte //3位长的 hms
	Lng       uint32 //经度 10e-6
	Lat       uint32 //纬度 10e-6
	Vec1      uint16 //卫星定位速度
	Vec2      uint16 //车载定位速度
	Vec3      uint32 //行驶里程
	Direction uint16 // 方向 北=>0
	Altitude  uint16 //海拔
	State     uint32 //车辆状态 忽略
	Alarm     uint32 //车辆告警 忽略
}

func (*GNSSData) Len() int {
	return 1 + 4 + 3 + 4 + 4 + 2 + 2 + 4 + 2 + 2 + 4 + 4
}

func (gnss *GNSSData) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Encrypt); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Date); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Time); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Lng); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Lat); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Vec1); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Vec2); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Vec3); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Direction); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Altitude); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.State); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &gnss.Alarm); err != nil {
		return err
	}
	return nil
}

func (gnss *GNSSData) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, gnss.Encrypt); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Date); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Time); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Lng); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Lat); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Vec1); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Vec2); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Vec3); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Direction); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Altitude); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.State); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, gnss.Alarm); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func NewGNSSData() GNSSData {
	return GNSSData{
		Date: make([]byte, 4, 4),
		Time: make([]byte, 3, 3),
	}
}
