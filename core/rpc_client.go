package core

import "github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
import log "github.com/alecthomas/log4go"

type RpcClient interface {
	init()
	start()
	invoke(service string, method string, msg *rpc.Msg) *rpc.Msg
}

type RpcClientImpl struct {
}

func (client *RpcClientImpl) init() {
	log.Info("rpcClient init.")
}
func (client *RpcClientImpl) start() {
	log.Info("rpcClient start.")
}
func (client *RpcClientImpl) invoke(service string, method string, msg *rpc.Msg) *rpc.Msg {
	return msg;
}
