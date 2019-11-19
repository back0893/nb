package net

import (
	"Nb/iface"
	"Nb/utils"
	"bufio"
	"fmt"
	"net"
)

type Connection struct {
	conn     *net.TCPConn
	connId   uint64
	ExitChan chan bool
	IsStop   bool
}

func NewConnection(connId uint64, conn *net.TCPConn) iface.IConnection {
	return &Connection{
		conn:     conn,
		connId:   connId,
		ExitChan: make(chan bool),
		IsStop:   false,
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
	go c.StartRead()
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
	for scan.Scan() {
		data := scan.Bytes()
		utils.LoggerObject.Write(string(data))
		c.Write(append(data, '\n'))
	}
}
func (c *Connection) Stop() {
	if c.IsStop {
		return
	}
	c.IsStop = true
	c.ExitChan <- true
	close(c.ExitChan)
	c.conn.Close()
}
