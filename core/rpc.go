package core

var app *ApplicationImpl

type Rpc struct {
}

func (rpc *Rpc) CreateApp() *ApplicationImpl {
	app = app.Init()
	return app
}
func (rpc *Rpc) GetAccessService() *RpcClientImpl {
	return app.GetRpcClient()
}

func (rpc *Rpc) GetRpcService() *RpcServerImpl {
	return app.GetRpcServer();
}
