package core

var app *ApplicationImpl

type Rpc struct {
}

func (rpc *Rpc) CreateApp(configPath string) *ApplicationImpl {
	//初始化应用
	app = app.Init(configPath)
	return app
}

func (rpc *Rpc) GetRpcService() *RpcClientImpl {
	return app.GetRpcClient()
}

func (rpc *Rpc) GetAccessService() *RpcServerImpl {
	return app.GetRpcServer();
}
