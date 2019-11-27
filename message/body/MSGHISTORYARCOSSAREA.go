package body

import (
	"bytes"
	"encoding/binary"
)

type HISTORYARCOSSAREA struct {
	CarNum    []byte //长21位的车牌 gbk
	Color     byte
	DataType  uint16 //子业务标识
	Length    uint32
	GNSSCnt   byte //后续包含的gnss_data的个数
	GNSSDatas []GNSSData
}

func (msg *HISTORYARCOSSAREA) UnmarshalUn(data []byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &msg.Length); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.GNSSCnt); err != nil {
		return err
	}
	for i := 0; i < int(msg.GNSSCnt); i++ {
		gnss := GNSSData{}
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
		msg.GNSSDatas = append(msg.GNSSDatas, gnss)
	}

	return nil
}

func (msg *HISTORYARCOSSAREA) Marshal() ([]byte, error) {
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
	if err := binary.Write(buffer, binary.BigEndian, &msg.Length); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &msg.GNSSCnt); err != nil {
		return nil, err
	}
	for i := 0; i < int(msg.GNSSCnt); i++ {
		gnss := msg.GNSSDatas[i]
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Date); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Time); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Lng); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Lat); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Vec1); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Vec2); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Vec3); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Direction); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Altitude); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.State); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &gnss.Alarm); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}
