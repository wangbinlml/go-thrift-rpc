package main

import (
	"flag"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"os"
	"os/signal"
	"syscall"
	"github.com/wangbinlml/go-thrift-rpc/example/biz/app/dispatcher"
	"runtime"
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	var configFile = "example/biz/config"

	//业务分发
	var iDis core.IBizDispatcher
	var biz = make(map[string]core.IBizDispatcher)
	iDis = new(dispatcher.BizDispatcher)

	//分发到不同业务
	biz["biz_service"] = iDis

	//创建RPC框架
	rpc := core.Rpc{}
	app := rpc.CreateApp(configFile, biz)

	//启动框架
	app.Start()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-sc
}