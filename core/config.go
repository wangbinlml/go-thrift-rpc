package core

import (
	io "io/ioutil"
)

type AcceptorConfig struct {
	Name    string `json:"name"`
	Service string `json:"service"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
}

type ConnectorConfig struct {
	Name          string `json:"name"`
	Service       string `json:"service"`
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
