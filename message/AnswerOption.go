package message

import (
	"bytes"
	"fmt"
	"strconv"
)

type AnswerOption struct {
	Record      int
	Upload      int
	VoltageRate int //电压波动
	VoltageMax  int //电压上限
	VoltageMin  int //电压下限
	rawData     []byte
}

func NewAnswerOption() *AnswerOption {
	return &AnswerOption{}
}
func (answer *AnswerOption) UnmarshalUn(data []byte) error {
	answer.rawData = data
	//A:30:01:P1000H4000L1000
	answer.rawData = data
	//去掉A: 在使用, 分割
	data = data[2:]

	splice := bytes.Split(data, []byte{':'})
	record, err := strconv.Atoi(string(splice[0]))
	if err != nil {
		return err
	}
	answer.Record = record

	upload, err := strconv.Atoi(string(splice[1]))
	if err != nil {
		return err
	}
	answer.Upload = upload

	voltage, err := strconv.Atoi(string(splice[2][1:5]))
	if err != nil {
		return err
	}

	answer.VoltageRate = voltage

	voltage, err = strconv.Atoi(string(splice[2][6:10]))
	if err != nil {
		return err
	}

	answer.VoltageMin = voltage

	voltage, err = strconv.Atoi(string(splice[2][11:]))
	if err != nil {
		return err
	}

	answer.VoltageMax = voltage

	return nil
}
func (answer *AnswerOption) GetId() uint32 {
	return 3
}

func (AnswerOption) Marshal() []byte {
	return []byte("{end}")
}

func (answer *AnswerOption) GetData() []byte {
	return answer.rawData
}
func (answer *AnswerOption) String() string {
	return fmt.Sprintf("配置%s:%d:%d:P%dH%dL%d", "S", answer.Record, answer.Upload, answer.VoltageMax, answer.VoltageMax, answer.VoltageMin)
}
