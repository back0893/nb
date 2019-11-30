package utils

import (
	"sync"
	"time"
)

type DeviceUid string

/**
一个定时检查是否上下线的
*/
type DeviceUpload struct {
	devices sync.Map
}
type ConCheck struct {
	Upload int   //上传次数
	Time   int64 //最新一次上传时间
	IsFist bool  //如果是第一次上线才回去更新gateway的is_online
}

func NewConCheck(upload int, time int64) *ConCheck {
	return &ConCheck{
		Upload: upload,
		Time:   time,
	}
}
func NewDeviceUpload() *DeviceUpload {
	return &DeviceUpload{}
}
func (upload *DeviceUpload) Store(key string, value *ConCheck) {
	upload.devices.Store(key, value)
}
func (upload *DeviceUpload) Load(key string) (*ConCheck, bool) {
	value, ok := upload.devices.Load(key)
	return value.(*ConCheck), ok
}
func (upload *DeviceUpload) Range() {
	fn := func(key interface{}, value interface{}) bool {
		device_id := key.(string)
		con_check := value.(*ConCheck)
		var intval int64
		if con_check.Upload == 99 {
			intval = 120
		} else {
			intval = int64(24 * 3600 / con_check.Upload)
		}
		timestamp := time.Now().Unix()
		sql := "update pra set is_online=? where deivce_id=?"
		var is_online int = 1
		if timestamp-con_check.Time > intval {
			//不在线
			is_online = 0
		}
		GlobalObject.Db.Exec(sql, is_online, device_id)
		return true
	}
	upload.devices.Range(fn)
}
