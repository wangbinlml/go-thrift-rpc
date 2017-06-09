package business

import "github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"

type BusinessService interface {
	doBusiness(messageName string, msg *rpc.Msg) (r *rpc.Msg, err error)
}
