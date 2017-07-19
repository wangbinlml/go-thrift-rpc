package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/logs"
)

var thriftAcceptor = new(ThriftAcceptor)
type RpcServer interface {
	init()
	initBiz()
	start()
}
type RpcServerImpl struct {
}

func (server *RpcServerImpl) init(acceptor AcceptorConfig, biz map[string] IBizDispatcher) {
	thriftAcceptor.init(acceptor, biz)
	logs.Info("rpcServer init.")
}
func (server *RpcServerImpl) start() {
	thriftAcceptor.start()
	logs.Info("rpcServer start.")
}
