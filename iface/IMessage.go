package iface

type IMessage interface {
	UnmarshalUn([]byte) error
	Marshal() ([]byte, error)
	GetData() []byte
	GetId() uint32
}
