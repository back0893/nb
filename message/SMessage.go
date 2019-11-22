package message

import (
	"bytes"
	"errors"
	"fmt"
)

/**
平台下发修改终端参数
*/
type SMessage struct {
	Record      int
	Upload      int
	VoltageRate float32 //电压波动
	VoltageMax  float32 //电压上限
	VoltageMin  float32 //电压下限
}

func (SMessage) UnmarshalUn([]byte) error {
	return errors.New("未实现")
}

func (msg *SMessage) Marshal() []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte('{')
	str := fmt.Sprintf("%s:%d:%d:P%.3fH%.3fL%.3f", "S", msg.Record, msg.Upload, msg.VoltageMax, msg.VoltageMax, msg.VoltageMin)
	buffer.WriteString(str)
	buffer.WriteByte(' ')
	buffer.WriteString(str)
	buffer.WriteByte('}')
	return buffer.Bytes()
}

func (SMessage) GetData() []byte {
	return nil
}

func (SMessage) GetId() uint32 {
	return 2
}
