package message

/**
平台回应的因为格式不一样....
依据格式回应不同的值,
mdzz
*/
type Answer struct {
}

func (Answer) UnmarshalUn([]byte) error {
	//终端回应的有2种
	//判断是否有,出现,有就是ip的回应
	//么有就是修改终端参数的命令

}

func (Answer) Marshal() []byte {
	panic("implement me")
}

func (Answer) GetData() []byte {
	panic("implement me")
}
