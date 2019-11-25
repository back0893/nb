package iface

/**
数据体
*/
type IBody interface {
	UnmarshalUn([]byte) error
	Marshal() ([]byte, error)
}
