package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kaidev1024/gokai/zinx/ziface"
)

// provide global config parameters
// some params can be customized by users in zinx.json

type GlobalObj struct {
	//server
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string
	//zinx
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}
