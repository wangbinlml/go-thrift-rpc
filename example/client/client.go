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
	"runtime"
	"time"
)

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
func SayHello(w http.ResponseWriter, req *http.Request) {
	startTime := currentTimeMillis()

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

	endTime := currentTimeMillis()
	fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))

	if error != nil {
		fmt.Println(error)
		w.Write([]byte("error"))
	} else {
		body := r1.GetBody()
		fmt.Println(r1.Header.Tid)
		w.Write([]byte(body))
	}

}
func main() {

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	//启动RPC
	go startRpc()

	//启动HTTP
	go startHttp()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-sc
}

func startHttp() {
	//传递处理请求方法
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)
	fmt.Println("start http server .....")
}

func startRpc() {
	var configFile = "example/client/config"
	var biz = make(map[string]core.IBizDispatcher)

	//创建RPC框架
	rpc := core.Rpc{}
	app := rpc.CreateApp(configFile, biz)

	//启动RPC框架
	app.Start()
}
