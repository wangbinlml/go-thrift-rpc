package core

import (
	io "io/ioutil"
	"github.com/tidwall/gjson"
)

var rpcConfig RpcConfig
var zkConfig ZKConfig

type AcceptorConfig struct {
	Name    string `json:"name"`
	Service string `json:"service"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Version    string `json:"version"`
	Weight    string `json:"weight"`
}

type ConnectorConfig struct {
	Name          string `json:"name"`
	Service       string `json:"service"`
	Version    string `json:"version"`
	RetryTime     string `json:"retryTime"`
	RetryInterval string `json:"retryInterval"`
	MaxPoolSize   string `json:"maxPoolSize"`
	MinPoolSize   string `json:"minPoolSize"`
	IdleTimeout   string `json:"idleTimeout"`
}

type RpcConfig struct {
	Acceptor  AcceptorConfig `json:"acceptor"`
	Connector []ConnectorConfig `json:"connector"`
}

type ZKConfigOption struct {
	Retries        int32 `json:"retries"`
	SessionTimeout int32 `json:"sessionTimeout"`
	SpinDelay      int32 `json:"spinDelay"`
}
type ZKConfig struct {
	Zk_path   string `json:"zk_path"`
	Zk_option ZKConfigOption `json:"zk_option"`
}


func Load(filename string) []byte {
	data, err := io.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	datajson := []byte(data)
	return datajson
}

func InitConfig(configPath string)  {
	data := Load(configPath + "/rpcServiceConfig.json")
	rpcConfig.Acceptor = AcceptorConfig{}
	rpcConfig.Connector = []ConnectorConfig{}
	gjson.Unmarshal(data, &rpcConfig)

	zkConfig.Zk_option = ZKConfigOption{}
	gjson.Unmarshal(Load(configPath + "/config.json"), &zkConfig)
}

func (config * RpcConfig) getRpcConfig() RpcConfig {
	return rpcConfig
}

func (config * ZKConfig) getZkConfig() ZKConfig {
	return zkConfig
}