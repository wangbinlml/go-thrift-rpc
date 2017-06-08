package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"fmt"
)

type RPCInvokeServiceImpl struct {
}

func (this *RPCInvokeServiceImpl) Invoke(serviceName string, methodName string, msg *rpc.Msg) (r *rpc.Msg, err error) {
	body := msg.GetBody()
	fmt.Println("serviceName: " + serviceName + "===Body: ",body)
	msg.Body = body+" world!!"
	return msg, nil
}
