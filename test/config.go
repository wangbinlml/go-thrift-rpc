package main

import "fmt"
import (
	"github.com/wangbinlml/go-thrift-rpc/core"
	"encoding/json"
	"reflect"
)

func main() {
	JsonParse := core.NewJsonStruct()
	//解析zk
	v := core.ZKConfig{}
	v.Zk_option = core.ZKConfigOption{}
	JsonParse.Load("config/config.json", &v)
	fmt.Println(v.Zk_path)
	v_option:= v.Zk_option
	fmt.Println(v_option.Retries)
	fmt.Println(v_option.SessionTimeout)
	fmt.Println(v_option.SpinDelay)

	//解析rpcConfig
	r:=core.RpcConfig{}

	c:= core.ConnectorConfig{}
	c.ConnectorService = []core.ConnectorConfigService{}

	r.Acceptor = core.AcceptorConfig{}
	r.Connector = c

	JsonParse.Load("config/rpcServiceConfig.json", &r)
	fmt.Println(r)


	var objs interface{}
	if err := json.Unmarshal(core.Load("config/rpcServiceConfig.json"), &objs); err != nil {
		fmt.Println(err)
		return
	}
	for _, obj := range objs {
		core.Switch(reflect.ValueOf(obj))
	}
}
