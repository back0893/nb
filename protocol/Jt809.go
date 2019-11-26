package protocol

import (
	"bytes"
	"io"
)

type Jt809 struct{}

func (Jt809) SplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 && atEOF == true {
		return 0, nil, io.EOF
	}
	start_index := bytes.IndexByte(data, 0x5b)
	if start_index == -1 {
		//灭有开始标志,跳过
		return len(data), nil, nil
	}
	end_index := bytes.IndexByte(data, 0x5d)
	if end_index == -1 && atEOF == true {
		//数据不全导致
		return 0, nil, io.EOF
	}
	if start_index >= end_index {
		//上一个包异常
		return start_index, nil, nil
	}
	message := data[start_index+1 : end_index]
	return end_index + 1, message, nil
}

func (Jt809) Decode(data []byte) []byte {
	//转义
	/**
	在数据发送时进行了如下的转义
	转义 0x5b=>0x5a 0x01
	转义 0x5a=>0x5a 0x02
	转义 0x5d=>0x5e 0x01
	转义 0x5e=>0x5e 0x01
	*/
	data = bytes.ReplaceAll(data, []byte{0x5a, 0x01}, []byte{0x5b})
	data = bytes.ReplaceAll(data, []byte{0x5a, 0x02}, []byte{0x5a})
	data = bytes.ReplaceAll(data, []byte{0x5e, 0x01}, []byte{0x5d})
	data = bytes.ReplaceAll(data, []byte{0x5e, 0x02}, []byte{0x5e})
	return data
}

func (Jt809) Encode(data []byte) []byte {
	buffer := bytes.NewBuffer([]byte{0x5b})
	data = bytes.ReplaceAll(data, []byte{0x5b}, []byte{0x5a, 0x01})
	data = bytes.ReplaceAll(data, []byte{0x5a}, []byte{0x5a, 0x02})
	data = bytes.ReplaceAll(data, []byte{0x5d}, []byte{0x5e, 0x01})
	data = bytes.ReplaceAll(data, []byte{0x5e}, []byte{0x5e, 0x02})
	buffer.Write(data)
	buffer.WriteByte(0x5d)
	return buffer.Bytes()
}
