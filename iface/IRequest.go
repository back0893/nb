package iface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetMsg() IMessage
	SetProperty(key string, value interface{})  //设置属性
	GetProperty(key string) (interface{}, bool) //读取
}
