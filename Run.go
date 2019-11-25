package main

import (
	"Nb/net"
	"Nb/router"
	"bytes"
	"io"
)

func main() {
	server := net.NewServer()
	server.AddSplitFunc(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
		if start_index <= end_index {
			//上一个包异常
			return start_index, nil, nil
		}
		message := data[start_index : end_index+1]
		/**
		在数据发送时进行了如下的转义
		转义 0x5b=>0x5a 0x01
		转义 0x5a=>0x5a 0x02
		转义 0x5d=>0x5e 0x01
		转义 0x5e=>0x5e 0x01
		*/
		message = bytes.ReplaceAll(message, []byte{0x5a, 0x01}, []byte{0x5b})
		message = bytes.ReplaceAll(message, []byte{0x5a, 0x02}, []byte{0x5a})
		message = bytes.ReplaceAll(message, []byte{0x5e, 0x01}, []byte{0x5d})
		message = bytes.ReplaceAll(message, []byte{0x5e, 0x02}, []byte{0x5e})
		return end_index + 1, message, nil

	})
	server.AddRouter(0x1001, router.NewHandler())
	server.Run()
}
