package core

var app Application

type Rpc struct {
}

func (rpc *Rpc) createApp() Application {
	return app
}
func GetAccessService() RpcClient {
	return app.GetRpcClient()
}

func GetRpcService() RpcServer {
	return app.GetRpcServer();
}
