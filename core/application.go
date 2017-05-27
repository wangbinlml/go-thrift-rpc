package core

import (
	log "github.com/alecthomas/log4go"
)
type ApplicationImpl struct {
}

func (app *ApplicationImpl) Init() *ApplicationImpl {
	rpcClient := RpcClientImpl{}
	rpcServer := RpcServerImpl{}
	rpcClient.init()
	rpcServer.init()
	log.Info("application init.")
	return app
}
func (app *ApplicationImpl) Start() {
	rpcClient := RpcClientImpl{}
	rpcServer := RpcServerImpl{}
	rpcClient.start()
	rpcServer.start()
	log.Info("application Start.")
}

func (app *ApplicationImpl) GetRpcClient() *RpcClientImpl {
	var rpcClient *RpcClientImpl

	return rpcClient
}
func (app *ApplicationImpl) GetRpcServer() *RpcServerImpl {
	var rpcServer *RpcServerImpl

	return rpcServer
}
