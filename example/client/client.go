package main

import (
	"flag"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"net/http"
	"fmt"
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

	r1, _ := client.Invoke("biz_service", "login", msg)
	body := r1.GetBody()
	w.Write([]byte(body))
}
func main() {
	flag.Parse()
	//go startRpc()
	startHttp()

}

func startHttp () {
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