package main

import (
	"flag"
	"github.com/wangbinlml/go-thrift-rpc/core"
)
func main() {
	flag.Parse()
	var configFile = "example/biz/config"
	rpc := core.Rpc{}
	app := rpc.CreateApp(configFile)
	app.Start()
}
