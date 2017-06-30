package business

import (
	"fmt"
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
)

type SMSBusiness struct {
}

func (biz *SMSBusiness) init() {
	fmt.Println("SMSBusiness init finished")
}

func (biz *SMSBusiness) DoBusiness(service string, method string, msg *rpc.Msg) (msg2 *rpc.Msg, err error) {
	fmt.Println("SMSBusiness doBusiness")
	/*rpc2 := core.Rpc{}
	client := rpc2.GetRpcService()
	client.Invoke()*/
	body := msg.GetBody()
	msg.Body = body + " world!!"
	return msg, nil
}
