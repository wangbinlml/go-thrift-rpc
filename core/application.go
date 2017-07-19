package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/logs"
)

type ApplicationImpl struct {
}

var connectorConfig []ConnectorConfig
var acceptorConfig AcceptorConfig

func (app *ApplicationImpl) Init(configPath string, biz map[string]IBizDispatcher) *ApplicationImpl {
	//初始化
	InitLog(configPath)
	//初始化配置项
	InitConfig(configPath)
	//初始化zk
	InitZk()

	var configInfo = new(RpcConfig)
	var rpcConfig = configInfo.getRpcConfig()
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

	logs.Info("application init.")

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
	logs.Info("application Start.")
}

func (app *ApplicationImpl) GetRpcClient() *RpcClientImpl {
	var rpcClient *RpcClientImpl
	return rpcClient
}
func (app *ApplicationImpl) GetRpcServer() *RpcServerImpl {
	var rpcServer *RpcServerImpl
	return rpcServer
}
