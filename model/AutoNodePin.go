package model

import (
	"time"
)

type AutoNodePing struct {
	ID       uint `gorm:"primary_key"`
	NodeId   int
	PinId    int
	Type     int
	PinName  string
	PinValue int
	Time     time.Time
	Unit     string
}

func (AutoNodePing) TableName() string {
	return "ywn_auto_node_pin"
}
