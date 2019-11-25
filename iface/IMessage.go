package iface

type IMessage interface {
	UnmarshalUn([]byte) error
	Marshal() []byte
	GetData() []byte
	GetId() uint32
	String() string
}
