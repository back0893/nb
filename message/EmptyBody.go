package message

import "Nb/iface"

type EmptyBody struct{}

func NewEmptyBody() iface.IBody {
	return &EmptyBody{}
}
func (EmptyBody) UnmarshalUn([]byte) error {
	return nil
}

func (EmptyBody) Marshal() ([]byte, error) {
	return []byte{}, nil
}

func (EmptyBody) Len() int {
	return 0
}
