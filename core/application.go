package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/logs"
)

type ApplicationImpl struct {
}

var connectorConfig []ConnectorConfig
var acceptorConfig AcceptorConfig

var rpcClient = RpcClientImpl{}
var rpcServer = RpcServerImpl{}

func (app *ApplicationImpl) Init(configPath string) *ApplicationImpl {
	//初始化
	InitLog(configPath)
	//初始化配置项
	InitConfig(configPath)
	//初始化zk
	InitZk()

	var configInfo = new(RpcConfig)
	var rpcConfig = configInfo.getRpcConfig()
	connectorConfig = rpcConfig.Connector
	acceptorConfig = rpcConfig.Acceptor
	if len(connectorConfig) > 0 {
		rpcClient.init(connectorConfig)
	}
	if acceptorConfig.Service != "" {
		rpcServer.init(rpcConfig.Acceptor)
	}

	logs.Info("application init.")

	return app
}

func (app *ApplicationImpl) RegistService(biz map[string]IBizDispatcher) {
	if acceptorConfig.Service != "" {
		rpcServer.registService(biz)
	}
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
