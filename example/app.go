package main

import (
	_ "github.com/wangbinlml/go-thrift-rpc/core"
	"flag"
	log "github.com/alecthomas/log4go"
	"os"
	"syscall"
	"github.com/wangbinlml/go-thrift-rpc/core"
)

func main() {
	flag.Parse()
	sc := make(chan os.Signal, 1)
	rpc := core.Rpc{}
	app := rpc.CreateApp()
	app.Start()

	sig := <-sc
	if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
		log.Info("main", "main", "Got signal", 0, "signal", sig)
		log.Close()
	} else if sig == syscall.SIGPIPE {
		log.Info("main", "main", "Ignore broken pipe signal", 0)
	}
}
