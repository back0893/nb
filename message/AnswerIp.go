package message

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/**
平台回应的因为格式不一样....
依据格式回应不同的值,
mdzz
*/
type AnswerIp struct {
	Ip      string
	Port    int
	rawData []byte
}

func (answer *AnswerIp) String() string {
	return fmt.Sprintf("返回的ip:%s:%d", answer.Ip, answer.Port)
}

func NewAnswerIp() *AnswerIp {
	return &AnswerIp{}
}
func (answer *AnswerIp) UnmarshalUn(data []byte) error {
	//A:"111.222.333.444","1234"
	answer.rawData = data
	//去掉A: 在使用, 分割
	data = data[2:]
	splice := bytes.Split(data, []byte{','})
	answer.Ip = string(splice[0])
	answer.Ip = strings.Trim(answer.Ip, "\"")
	p := strings.Trim(string(splice[1]), "\"")
	port, err := strconv.Atoi(p)
	if err != nil {
		return err
	}
	answer.Port = port
	return nil
}

func (answer *AnswerIp) GetId() uint32 {
	return 2
}

func (AnswerIp) Marshal() []byte {
	return []byte("{end}")
}

func (answer *AnswerIp) GetData() []byte {
	return answer.rawData
}
