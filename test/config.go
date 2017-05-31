package main

import (
	"github.com/tidwall/gjson"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"fmt"
	"reflect"
)

func main() {
	//解析zk
	fmt.Println("======解析zk=====")
	v := core.ZKConfig{}
	v.Zk_option = core.ZKConfigOption{}
	gjson.Unmarshal(core.Load("config/config.json"), &v)
	fmt.Println(v.Zk_path)
	v_option := v.Zk_option
	fmt.Println(v_option.Retries)
	fmt.Println(v_option.SessionTimeout)
	fmt.Println(v_option.SpinDelay)

	data := core.Load("config/rpcServiceConfig.json")
	config := string(data)
	result := gjson.Get(config, "connector")
	for _, v := range result.Array() {
		fmt.Println(v.Get("name"))
		fmt.Println("type:", reflect.TypeOf(v))
	}
	//解析rpcConfig
	fmt.Println("=====解析rpcConfig====")
	r := core.RpcConfig{}
	r.Acceptor = core.AcceptorConfig{}
	r.Connector = []core.ConnectorConfig{}
	gjson.Unmarshal(data, &r)
	fmt.Println(r)
}
