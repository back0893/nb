package body

type ConnectRsp struct {
	Result     byte
	VerifyCode uint32
}

func (ConnectRsp) UnmarshalUn([]byte) error {

}

func (ConnectRsp) Marshal() ([]byte, error) {
}
