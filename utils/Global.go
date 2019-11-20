package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var GlobalObject *Global

func init() {
	GlobalObject = &Global{
		Host: "127.0.0.1",
		Port: 8001,
	}
	GlobalObject.Reload()
}

type Global struct {
	Host string `json:"host"`
	Port int    `port:"port"`
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
