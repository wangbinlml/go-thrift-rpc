package main

import (
	"flag"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"os"
	"os/signal"
	"syscall"
	"github.com/wangbinlml/go-thrift-rpc/example/biz/app/dispatcher"
)

func main() {
	flag.Parse()
	//runtime.GOMAXPROCS(runtime.NumCPU())
	var configFile = "example/biz/config"
	var iDis core.IBizDispatcher
	var biz = make(map[string]core.IBizDispatcher)
	iDis = new(dispatcher.BizDispatcher)
	biz["biz_service"] = iDis
	rpc := core.Rpc{}
	app := rpc.CreateApp(configFile, biz)
	app.Start()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-sc
}
func router(make (map[string]interface{})) {

}
