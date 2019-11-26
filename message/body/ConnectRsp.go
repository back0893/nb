package body

type ConnectRsp struct {
	Result     byte
	VerifyCode uint32
}

func (ConnectRsp) UnmarshalUn([]byte) error {
	return nil
}

func (ConnectRsp) Marshal() ([]byte, error) {
	return []byte{}, nil
}
