package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"fmt"
)

type IBizDispatcher interface {
	Dispatch(service string, method string, msg *rpc.Msg) (msg2 *rpc.Msg, err error)
}

type RPCInvokeServiceImpl struct {
	service map[string]IBizDispatcher
}

func (rpc *RPCInvokeServiceImpl) Invoke(serviceName string, methodName string, msg *rpc.Msg) (r *rpc.Msg, err error) {
	fmt.Println("serviceName: " + serviceName)
	var dis IBizDispatcher
	for k, v := range rpc.service {
		if k == serviceName {
			dis = v
		}
	}
	return dis.Dispatch(serviceName, methodName, msg)
}
