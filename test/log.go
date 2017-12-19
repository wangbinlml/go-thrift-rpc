package main

import (
	"github.com/wangbinlml/go-thrift-rpc/core"
	"github.com/astaxie/beego/logs"
)

func main() {
	core.InitLog("test/")
	logs.Info("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.Info("this is a message of orm")

	logs.Alert("alert")
	logs.Notice("notice")
	logs.Emergency("Emergency")
	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
}
