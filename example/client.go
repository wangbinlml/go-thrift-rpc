package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"os"
	"time"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRPCInvokeServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	startTime := currentTimeMillis()
	for i := 0; i < 10000000; i++ {

		header := &rpc.Header{
		}
		header.Protocol = "thrift"
		header.Tid = "3232b32b32h3h2b32b3n2b3nb2n32"
		msg := &rpc.Msg{
		}
		msg.Header = header
		msg.Body = "hello"

		r1, e1 := client.Invoke("biz_service", "login", msg)
		body := r1.GetBody()
		fmt.Println(i, "Call->", body, e1)
	}

	endTime := currentTimeMillis()
	fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))
}

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
