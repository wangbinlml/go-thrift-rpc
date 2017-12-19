package core

import (
	"github.com/astaxie/beego/logs"
)

var thriftAcceptor = new(ThriftAcceptor)

type RpcServer interface {
	init()
	initBiz()
	start()
}
type RpcServerImpl struct {
}

func (server *RpcServerImpl) init(acceptor AcceptorConfig) {
	thriftAcceptor.init(acceptor)
	logs.Info("rpcServer init.")
}

func (server *RpcServerImpl) registService(biz map[string]IBizDispatcher) {
	thriftAcceptor.registService(biz)
}

func (server *RpcServerImpl) start() {
	thriftAcceptor.start()
	logs.Info("rpcServer start.")
}
