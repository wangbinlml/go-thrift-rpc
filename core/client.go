package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"github.com/astaxie/beego/logs"
)

type RpcClient interface {
	init()
	start()
	invoke(service string, method string, msg *rpc.Msg) *rpc.Msg
}

var thriftConnector = new(ThriftConnector)

type RpcClientImpl struct {
}

func (client *RpcClientImpl) init(connector []ConnectorConfig) {
	thriftConnector.init(connector)
	logs.Info("rpcClient init.")
}
func (client *RpcClientImpl) start() {
	thriftConnector.start()
	logs.Info("rpcClient start.")
}
func (client *RpcClientImpl) Invoke(service string, method string, msg *rpc.Msg) (*rpc.Msg, error) {
	return thriftConnector.Invoke(service, method, msg)
}
