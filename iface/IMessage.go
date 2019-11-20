package iface

type IMessage interface {
	UnmarshalUn([]byte) error
	Marshal() []byte
	GetData() []byte
}
