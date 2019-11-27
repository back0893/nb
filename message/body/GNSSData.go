package body

type GNSSData struct {
	Encrypt   byte   //是否加密
	Date      []byte //4位长的 dmyy
	Time      []byte //3位长的 hms
	Lng       uint32 //经度 10e-6
	Lat       uint32 //纬度 10e-6
	Vec1      uint16 //卫星定位速度
	Vec2      uint16 //车载定位速度
	Vec3      uint32 //行驶里程
	Direction uint16 // 方向 北=>0
	Altitude  uint16 //海拔
	State     uint32 //车辆状态 忽略
	Alarm     uint32 //车辆告警 忽略
}
