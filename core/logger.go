package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/logs"
	"log"
)
type Logger struct {
}

func int() {
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile,`{"filename":"/tmp/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
}

func (log *Logger) getLogger(category string) *log.Logger {
	return logs.GetLogger(category)
}
