package main

import (
	"flag"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"os"
	"os/signal"
	"syscall"
)
func main() {
	flag.Parse()

	var configFile = "example/biz/config"
	rpc := core.Rpc{}
	app := rpc.CreateApp(configFile)
	app.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-sc
}
