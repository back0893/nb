package EXGMSG

type EXGMSG struct {
	CarNum   []byte //长21位的车牌 gbk
	Color    byte
	DataType uint16 //子业务标识
	Length   uint32
}

func NewEXGMSG() EXGMSG {
	return EXGMSG{
		CarNum: make([]byte, 21, 21),
	}
}
