package message

import (
	"bytes"
)

/**
平台回应的因为格式不一样....
依据格式回应不同的值,
mdzz
*/
type Answer struct {
	rawData []byte
}

func (answer *Answer) GetId() uint32 {
	return 3
}

func (answer *Answer) UnmarshalUn(data []byte) error {
	answer.rawData = data
	//终端回应的有2种
	//判断是否有,出现,有就是ip的回应
	//么有就是修改终端参数的命令
	if bytes.IndexByte(data, ',') != -1 {
		//是ip回应
	} else {
		//是修改参数命令
	}
	return nil
}

func (Answer) Marshal() []byte {
	panic("implement me")
}

func (Answer) GetData() []byte {
	panic("implement me")
}
