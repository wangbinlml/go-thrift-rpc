package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"github.com/wangbinlml/go-thrift-rpc/core"
)
const (
	NetworkAddr = "127.0.0.1:19090"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &core.RPCInvokeServiceImpl{}
	processor := rpc.NewRPCInvokeServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
