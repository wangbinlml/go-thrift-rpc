package core

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"fmt"
	zookeeper "github.com/samuel/go-zookeeper/zk"
	"os"
	log "github.com/alecthomas/log4go"
)

var zkUtil = new(ZKUtil)

type ThriftAcceptor struct {
	Host    string
	Port    string
	Name    string
	Version string
	Weight  string
	ZkPath  string
}

func (acceptor *ThriftAcceptor) init(ac AcceptorConfig) {
	acceptor.Host = ac.Ip
	acceptor.Port = ac.Port
	acceptor.Name = ac.Name
	if ac.Version == "" {
		acceptor.Version = "1.0.0"
	} else {
		acceptor.Version = ac.Version
	}
	if ac.Weight == "" {
		acceptor.Weight = "1"
	} else {
		acceptor.Weight = ac.Weight
	}
	acceptor.ZkPath = ac.Service
	log.Info("ThriftAcceptor init")
}

func (acceptor *ThriftAcceptor) start() {
	var networkAddr = acceptor.Host + ":" + acceptor.Port
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(networkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RPCInvokeServiceImpl{}
	processor := rpc.NewRPCInvokeServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	log.Info("thrift server in " + networkAddr)

	//注册zk服务
	path := acceptor.ZkPath + "/" + acceptor.Version
	flags := int32(zookeeper.FlagEphemeral)

	err = zkUtil.Create(path, "", int32(0))
	if err != nil {
		fmt.Printf("create: %+v\n", err)
	}
	path = path + "/" + acceptor.Host + ":" + acceptor.Port + ":" + acceptor.Weight;
	err = zkUtil.Create(path, "", flags)
	Must(err)
	log.Info("ThriftAcceptor start " + path)
	server.Serve()
}
