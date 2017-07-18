package main

import (
	"github.com/wangbinlml/go-thrift-rpc/core/logs"
	"github.com/wangbinlml/go-thrift-rpc/core"
)

func main() {
	filename := "/home/wangbin/workspace/golang/src/github.com/wangbinlml/go-thrift-rpc/config/log4go.json"
	config := string(core.LoadFile(filename))
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, config)
	logs.SetLevel(7)
	//日志默认不输出调用的文件名和文件行号
	logs.EnableFuncCallDepth(true)
	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
}
