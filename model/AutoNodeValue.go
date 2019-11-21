package model

import (
	"time"
)

type AutoNodeValue struct {
	ID     uint `gorm:"primary_key"`
	NodeId int
	Pin    int
	Type   int
	Value  int
	Time   time.Time
}

func (AutoNodeValue) TableName() string {
	return "ywn_auto_node_value"
}
