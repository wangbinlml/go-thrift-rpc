package core

type Application struct {
}

func (app *Application) Init() *Application {
	
	return app
}
func (app *Application) Start() {

}

func (app *Application) GetRpcClient() *RpcClient {
	var rpcClient *RpcClient
	return rpcClient
}
func (app *Application) GetRpcServer() *RpcServer {
	var rpcServer *RpcServer
	return rpcServer
}
