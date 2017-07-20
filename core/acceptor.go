package core

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	zookeeper "github.com/samuel/go-zookeeper/zk"
	"os"
	"github.com/wangbinlml/go-thrift-rpc/core/logs"
)

var zkUtil = new(ZKUtil)

type ThriftAcceptor struct {
	Host    string
	Port    string
	Name    string
	Version string
	Weight  string
	ZkPath  string
	biz     map[string]IBizDispatcher
}

func (acceptor *ThriftAcceptor) initBiz() {
	logs.Info("ThriftAcceptor initBiz")
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
	logs.Info("ThriftAcceptor init")
}

//注册业务处理服务
func (acceptor *ThriftAcceptor) registService(biz map[string]IBizDispatcher) {
	acceptor.biz = biz
	acceptor.initBiz()
	logs.Info("Regist server ",biz)
}

func (acceptor *ThriftAcceptor) start() {
	var networkAddr = acceptor.Host + ":" + acceptor.Port
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(networkAddr)
	if err != nil {
		logs.Error("Error!", err)
		os.Exit(1)
	}

	handler := &RPCInvokeServiceImpl{acceptor.biz}
	processor := rpc.NewRPCInvokeServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	logs.Info("thrift server in " + networkAddr)

	//注册zk服务
	path := acceptor.ZkPath + "/" + acceptor.Version
	flags := int32(zookeeper.FlagEphemeral)

	err = zkUtil.Create(path, "", int32(0))
	if err != nil {
		logs.Error("create: %+v\n", err)
	}
	path = path + "/" + acceptor.Host + ":" + acceptor.Port + ":" + acceptor.Weight;
	err = zkUtil.Create(path, "", flags)
	Must(err)
	logs.Info("start thrift rpc listen addr: %s", path)
	go server.Serve()
}
