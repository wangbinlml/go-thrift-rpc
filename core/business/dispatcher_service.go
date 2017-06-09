package business

import "github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"

type DispatchService interface {
	Dispatch(msg *rpc.Msg) (r *rpc.Msg, err error)
	Init(businessMap []interface{})
}