package net

import (
	"Nb/iface"
	"Nb/message"
	"Nb/utils"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	conn        *net.TCPConn
	connId      uint64
	ExitChan    chan bool
	IsStop      bool
	server      iface.IServer
	msgChan     chan []byte
	msgBuffChan chan []byte
}

func NewConnection(connId uint64, conn *net.TCPConn, server iface.IServer) iface.IConnection {
	return &Connection{
		conn:        conn,
		connId:      connId,
		ExitChan:    make(chan bool),
		IsStop:      false,
		server:      server,
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, 1024),
	}
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.conn
}

func (c *Connection) Write(data []byte) (int, error) {
	return c.conn.Write(data)
}
func (c *Connection) GetConId() uint64 {
	return c.connId
}
func (c *Connection) Start() {
	c.server.CallOnConnStart(c)
	go c.StartRead()
	go c.StartWrite()
	for {
		select {
		case <-c.ExitChan:
			{
				return
			}
		}
	}
}
func (c *Connection) StartRead() {
	defer c.Stop()
	defer utils.LoggerObject.Write(fmt.Sprintf("%d连接退出", c.connId))
	scan := bufio.NewScanner(c.conn)
	scan.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if len(data) == 0 && atEOF == true {
			return 0, nil, io.EOF
		}
		start_index := bytes.IndexByte(data, '{')
		end_index := bytes.IndexByte(data, '}')
		if atEOF == true && (start_index == -1 || end_index == -1) {
			return 0, nil, io.EOF
		}
		if start_index == -1 || end_index == -1 {
			return 0, nil, nil
		}
		if start_index > end_index {
			//异常的流,寻找下一个争取的包
			return end_index + 1, nil, nil
		}
		return end_index + 1, data[start_index+1 : end_index], nil
	})
	for scan.Scan() {
		data := scan.Bytes()
		msg := message.NewMessage()
		err := msg.UnmarshalUn(data)
		fmt.Println(err)
		request := NewRequest(c, msg)
		go c.server.GetMsgRouter().DoMsgHandler(request)
	}
	<-c.ExitChan
}
func (c *Connection) StartWrite() {
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.conn.Write(data); err != nil {
				utils.LoggerObject.Write(err.Error())
			}
		case data := <-c.msgBuffChan:
			if _, err := c.conn.Write(data); err != nil {
				utils.LoggerObject.Write(err.Error())
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) SendMsg(data []byte) {
	c.msgChan <- data
}
func (c *Connection) SendBuffMsg(data []byte) {
	c.msgBuffChan <- data
}
func (c *Connection) Stop() {
	if c.IsStop {
		return
	}
	c.IsStop = true
	c.ExitChan <- true
	close(c.ExitChan)
	c.server.GetManager().Remove(c.connId)
	c.server.CallOnConnStop(c)
	c.conn.Close()
}
