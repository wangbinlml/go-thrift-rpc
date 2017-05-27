package core

import (
	io "io/ioutil"
	"encoding/json"
)

type AcceptorConfig struct {
	Name    string
	Service string
	Ip      string
	Port    string
}

type ConnectorConfigService struct {
	Name          string
	Service       string
	RetryTime     string
	RetryInterval string
	MaxPoolSize   string
	MinPoolSize   string
	IdleTimeout   string
}
type ConnectorConfig struct {
	ConnectorService []ConnectorConfigService `json:"connector"`
}

type RpcConfig struct {
	Acceptor  AcceptorConfig
	Connector ConnectorConfig
}

type ZKConfigOption struct {
	Retries        int32
	SessionTimeout int32
	SpinDelay      int32
}
type ZKConfig struct {
	Zk_path   string
	Zk_option ZKConfigOption
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {

	return &JsonStruct{}

}
func Load(filename string) []byte {
	data, err := io.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	datajson := []byte(data)
	return datajson
}
func (self *JsonStruct) Load(filename string, v interface{}) {

	data, err := io.ReadFile(filename)

	if err != nil {
		panic(err)
		return
	}

	datajson := []byte(data)

	err = json.Unmarshal(datajson, v)

	if err != nil {

		return

	}

}
