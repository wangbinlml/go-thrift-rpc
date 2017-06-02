package core

import (
	log "github.com/alecthomas/log4go"
)

var thriftAcceptor = new(ThriftAcceptor)
type RpcServer interface {
	init()
	start()
}
type RpcServerImpl struct {
}

func (server *RpcServerImpl) init(acceptor AcceptorConfig) {
	thriftAcceptor.init(acceptor)
	log.Info("rpcServer init.")
}
func (server *RpcServerImpl) start() {
	thriftAcceptor.start()
	log.Info("rpcServer start.")
}
