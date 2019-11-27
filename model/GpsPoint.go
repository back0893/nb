package model

type GpsPoint struct {
	ID         uint `gorm:"primary_key"`
	PlateNum   string
	PointX     float32
	PointY     float32
	VehicleId  uint
	AreaPolyId uint
	Time       uint
	Height     uint
	Speed      float32
	Direction  uint
	SimCard    string
}
