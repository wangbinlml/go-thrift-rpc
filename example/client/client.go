package main

import (
	"flag"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"net/http"
	"fmt"
	"os"
	"syscall"
	log "github.com/alecthomas/log4go"
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
	if error!=nil {
		fmt.Println(error)
		w.Write([]byte("error"))
	}else {
		body := r1.GetBody()
		w.Write([]byte(body))
	}
}
func main() {
	flag.Parse()
	sc := make(chan os.Signal, 1)
	var configFile = "example/client/config"
	rpc3 := core.Rpc{}
	app := rpc3.CreateApp(configFile)
	app.Start()

	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)
	fmt.Println("start http server .....")

	sig := <-sc
	if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
		log.Info("main", "main", "Got signal", 0, "signal", sig)
		log.Close()
	} else if sig == syscall.SIGPIPE {
		log.Info("main", "main", "Ignore broken pipe signal", 0)
	}
	/*fmt.Println("==========")
	time.Sleep(10000)
	header := &rpc.Header{
	}
	header.Protocol = "thrift"
	header.Tid = "3232b32b32h3h2b32b3n2b3nb2n32"
	msg := &rpc.Msg{
	}
	msg.Header = header
	msg.Body = "hello"

	client := rpc3.GetRpcService()
	r1, _ := client.Invoke("biz_service", "login", msg)
	body := r1.GetBody()
	fmt.Println(body)*/
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