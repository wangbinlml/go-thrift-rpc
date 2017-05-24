package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
)
const (
	NetworkAddr = "127.0.0.1:19090"
)

type RRPCInvokeServiceImpl struct {
}

func (this *RRPCInvokeServiceImpl) Invoke(serviceName string, methodName string, msg *core.Msg) (r *core.Msg, err error) {
	fmt.Println("-->FunCall:", msg)

	return
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RRPCInvokeServiceImpl{}
	processor := rpc.NewRPCInvokeServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
