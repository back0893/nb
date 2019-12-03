package utils

import (
	"Nb/iface"
	"encoding/json"
	"github.com/axgle/mahonia"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
)

var GlobalObject *Global

func init() {
	GlobalObject = &Global{
		Host:               "127.0.0.1",
		Port:               8001,
		MaxWorkerSize:      4,
		MaxWorkerQueueTask: 1024,
	}
	GlobalObject.Reload()
}

type DB struct {
	User   string            `json:"username"`
	Passwd string            `json:"password"`
	Net    string            `json:"net"`
	Addr   string            `json:"addr"`
	DbName string            `json:"dbName"`
	Params map[string]string `json:"params"`
}
type Global struct {
	Host               string         `json:"host"`
	Port               int            `json:"port"`
	Database           map[string]*DB `json:"database"`
	Server             iface.IServer
	Db                 map[string]*gorm.DB
	MaxWorkerSize      int `json:"worker_size"`
	MaxWorkerQueueTask int `json:"queue_task"`
}

func (global *Global) Reload() {
	fp, err := os.Open("./server.json")
	if err != nil {
		LoggerObject.Write(err.Error())
	}
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		LoggerObject.Write(err.Error())
	}
	if err := json.Unmarshal(data, GlobalObject); err != nil {
		LoggerObject.Write(err.Error())
	}
}

func (global *Global) ConvertToString(src []byte, srcCode string, tagCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(string(src))
	tagCoder := mahonia.NewEncoder(tagCode)
	cdata := tagCoder.ConvertString(srcResult)
	return []byte(cdata)
}
