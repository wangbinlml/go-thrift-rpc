package core

import (
	log "github.com/alecthomas/log4go"
)

type ApplicationImpl struct {
}

var connectorConfig []ConnectorConfig
var acceptorConfig AcceptorConfig

func (app *ApplicationImpl) Init(biz map[string]IBizDispatcher) *ApplicationImpl {
	var r = new(RpcConfig)
	var rpcConfig = r.getRpcConfig()
	rpcClient := RpcClientImpl{}
	rpcServer := RpcServerImpl{}
	connectorConfig = rpcConfig.Connector
	acceptorConfig = rpcConfig.Acceptor
	if len(connectorConfig) > 0 {
		rpcClient.init(connectorConfig)
	}
	if acceptorConfig.Service != "" {
		rpcServer.init(rpcConfig.Acceptor, biz)
	}

	log.Info("application init.")
	return app
}
func (app *ApplicationImpl) Start() {
	rpcClient := RpcClientImpl{}
	rpcServer := RpcServerImpl{}
	if len(connectorConfig) > 0 {
		rpcClient.start()
	}
	if acceptorConfig.Service != "" {
		rpcServer.start()
	}
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
