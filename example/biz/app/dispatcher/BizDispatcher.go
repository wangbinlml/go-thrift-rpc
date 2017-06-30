package dispatcher

import (
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	"fmt"
	"github.com/wangbinlml/go-thrift-rpc/example/biz/app/business"
)

type BizDispatcher struct {
}

func (dis *BizDispatcher) Dispatch(service string, method string, msg *rpc.Msg) (msg2 *rpc.Msg, err error){
	fmt.Println("SMSBusiness init dispatcher, service:",service)
	var biz = new (business.SMSBusiness)
	return biz.DoBusiness(service, method,msg)
}