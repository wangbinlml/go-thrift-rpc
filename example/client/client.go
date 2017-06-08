package main

import (
	"flag"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"net/http"
	"fmt"
	"os"
	"syscall"
	"os/signal"
)

func SayHello(w http.ResponseWriter, req *http.Request) {

	header := &rpc.Header{
	}
	header.Protocol = "thrift"
	header.Tid = "3232b32b32h3h2b32b3n2b3nb2n32"
	msg := &rpc.Msg{
	}
	msg.Header = header
	msg.Body = "hello"

	rpc2 := core.Rpc{}
	client := rpc2.GetRpcService()
	r1, error := client.Invoke("biz_service", "login", msg)
	if error != nil {
		fmt.Println(error)
		w.Write([]byte("error"))
	} else {
		body := r1.GetBody()
		w.Write([]byte(body))
	}
}
func main() {

	flag.Parse()

	go startRpc()
	go startHttp()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-sc
}

func startHttp() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)
	fmt.Println("start http server .....")
}

func startRpc() {
	var configFile = "example/client/config"
	rpc := core.Rpc{}
	app := rpc.CreateApp(configFile)
	app.Start()
}
