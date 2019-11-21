package model

type AutoNode struct {
	ID       uint `gorm:"primary_key"`
	Duid     string
	IsOnline string
}

func (AutoNode) TableName() string {
	return "ywn_auto_node"
}
