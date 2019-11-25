package message

import (
	"bytes"
	"fmt"
)

type IpMessage struct {
	Ip   string
	Port int
}

func (msg IpMessage) String() string {
	return fmt.Sprintf("I:\"%s\":%d", msg.Ip, msg.Port)
}

func (IpMessage) UnmarshalUn([]byte) error {
	panic("未实现")
}

func (msg *IpMessage) Marshal() []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte('{')
	str := msg.String()
	buffer.WriteString(str)
	buffer.WriteByte(' ')
	buffer.WriteString(str)
	buffer.WriteByte('}')
	return buffer.Bytes()
}

func (IpMessage) GetData() []byte {
	panic("未实现")
}

func (IpMessage) GetId() uint32 {
	return 5
}
