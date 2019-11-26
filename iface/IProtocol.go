package iface

type IProtocol interface {
	//分割
	SplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error)
	//进入时的转义
	Decode([]byte) []byte
	//发送时的转义
	Encode([]byte) []byte
}
