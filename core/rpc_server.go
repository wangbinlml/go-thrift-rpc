package core

import (
	log "github.com/alecthomas/log4go"
)

type RpcServer interface {
	init()
	start()
}
type RpcServerImpl struct {
}

func (server *RpcServerImpl) init() {
	log.Info("rpcServer init.")
}
func (server *RpcServerImpl) start() {
	log.Info("rpcServer start.")
}
