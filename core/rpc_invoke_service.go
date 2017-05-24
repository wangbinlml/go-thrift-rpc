package core

import "fmt"

type RPCInvokeServiceImpl struct {

}

func (this *RPCInvokeServiceImpl) Invoke(serviceName string, methodName string, msg *Msg) (r *Msg, err error) {
	fmt.Println("serviceName: " + serviceName)
	return msg, nil
}
