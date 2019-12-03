package message

import (
	"Nb/iface"
	"bytes"
	"encoding/binary"
	"github.com/howeyc/crc16"
)

/**
header 开始
*/
type Header struct {
	Len         uint32
	SN          uint32
	ID          uint16
	UUId        uint32
	Version     []byte
	EncryptFlag byte
	EncryptKey  uint32
}

func MakeHeader(id uint16, version []byte) *Header {
	header := NewHeader()
	header.ID = id
	header.Version = version
	return header
}

func NewHeader() *Header {
	return &Header{
		Version: make([]byte, 3, 3),
	}
}
func (header *Header) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &header.Len); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &header.SN); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &header.ID); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &header.UUId); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &header.Version); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &header.EncryptFlag); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &header.EncryptKey); err != nil {
		return err
	}
	return nil
}

func (header *Header) Marshal() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, &header.Len); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &header.SN); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &header.ID); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &header.UUId); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &header.Version); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &header.EncryptFlag); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &header.EncryptKey); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

/**
header 结束
*/

type Message struct {
	Header  *Header
	Body    iface.IBody
	Crc     uint16
	rawData []byte
}

func (msg *Message) UnmarshalUn(data []byte) error {
	buffer := bytes.NewBuffer(data)

	if err := msg.Header.UnmarshalUn(buffer.Next(22)); err != nil {
		return err
	}
	body_len := msg.Header.Len - 22 - 2 - 2
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
	body_data, err := msg.Body.Marshal()
	if err != nil {
		return nil, err
	}

	msg.Header.Len = 22 + 2 + 2 + uint32(msg.Body.Len())

	header_data, err := msg.Header.Marshal()
	if err != nil {
		return nil, err
	}
	buffer.Write(header_data)
	buffer.Write(body_data)
	//设置crc
	msg.Crc = crc16.ChecksumCCITTFalse(buffer.Bytes())
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

func (msg *Message) CheckSum() uint16 {
	data, _ := msg.Header.Marshal()
	b, _ := msg.Body.Marshal()
	data = append(data, b...)
	return crc16.ChecksumCCITTFalse(data)
}
