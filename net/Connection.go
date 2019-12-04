package net

import (
	"Nb/iface"
	"Nb/message"
	"Nb/message/body"
	"Nb/utils"
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

type Connection struct {
	conn        *net.TCPConn
	connId      uint64
	ExitChan    chan bool
	IsStop      bool
	server      iface.IServer
	msgChan     chan iface.IMessage
	msgBuffChan chan iface.IMessage
	lock        sync.RWMutex
	property    map[string]interface{}
}

func (c *Connection) RemoveProperty(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.property, key)
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.property[key] = value

}

func (c *Connection) GetProperty(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	t, ok := c.property[key]
	return t, ok
}

func NewConnection(connId uint64, conn *net.TCPConn, server iface.IServer) iface.IConnection {
	return &Connection{
		conn:        conn,
		connId:      connId,
		ExitChan:    make(chan bool),
		IsStop:      false,
		server:      server,
		msgChan:     make(chan iface.IMessage),
		msgBuffChan: make(chan iface.IMessage, 1024),
		property:    make(map[string]interface{}),
	}
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.conn
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
func (c *Connection) getSn() uint32 {

	value, ok := c.GetProperty("sn")
	var sn uint32
	defer c.SetProperty("sn", sn)

	if ok == true {
		sn := value.(uint32)
		sn++
	} else {
		sn = 1
	}
	return sn
}
func (c *Connection) recordTime() {
	timestamp := time.Now().Unix()
	c.SetProperty("timestamp", timestamp)
}
func (c *Connection) StartRead() {
	defer c.Stop()
	defer utils.LoggerObject.Write(fmt.Sprintf("%d连接退出", c.connId))

	scan := bufio.NewScanner(c.conn)

	scan.Split(c.server.GetProtocol().SplitFunc)
	for scan.Scan() {
		//停止就退出
		if c.IsStop {
			break
		}
		c.recordTime()
		//记录sn
		sn := c.getSn()
		//记录通信的时间,定时器检查

		data := scan.Bytes()
		msg, err := c.server.GetProtocol().Decode(data)
		if err != nil {
			utils.LoggerObject.Write(err.Error())
			continue
		}

		//判断是否登录成功
		if !c.checkLogin(msg) {
			continue
		}

		request := NewRequest(c, msg)
		request.SetProperty("sn", sn)
		if utils.GlobalObject.MaxWorkerSize > 0 {
			c.server.GetMsgRouter().SendMsgToTaskQueue(request)
		} else {
			go c.server.GetMsgRouter().DoMsgHandler(request)
		}
	}
	if !c.IsStop {
		c.ExitChan <- true
	}
}
func (c *Connection) checkLogin(msg iface.IMessage) bool {
	if msg.GetId() != 0x1001 {
		_, ok := c.GetProperty("login")
		if ok == false {
			//返回重新登录
			header := message.MakeHeader(0x1002, []byte{0x01, 0x00, 0x00})
			if value, ok := c.GetProperty("sn"); ok {
				header.SN = value.(uint32)
			} else {
				header.SN = 1
			}

			//1002的body回应
			body_msg := body.NewConnectRsp()
			body_msg.VerifyCode = 1
			body_msg.Result = 0x06

			response := &message.Message{
				Header: header,
				Body:   body_msg,
			}
			c.SendBuffMsg(response)
			return false
		}
	}
	return true
}
func (c *Connection) StartWrite() {
	for {
		select {
		case data := <-c.msgChan:
			//这里可以最终的转义处理,新增头尾标志
			raw, err := c.server.GetProtocol().Encode(data)
			if err != nil {
				utils.LoggerObject.Write(err.Error())
				return
			}
			if _, err := c.conn.Write(raw); err != nil {
				utils.LoggerObject.Write(err.Error())
			}
		case data := <-c.msgBuffChan:
			//这里可以最终的转义处理,新增头尾标志
			raw, err := c.server.GetProtocol().Encode(data)
			if err != nil {
				utils.LoggerObject.Write(err.Error())
				return
			}
			if _, err := c.conn.Write(raw); err != nil {
				utils.LoggerObject.Write(err.Error())
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) SendMsg(msg iface.IMessage) {
	c.msgChan <- msg
}
func (c *Connection) SendBuffMsg(msg iface.IMessage) {
	c.msgBuffChan <- msg
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
