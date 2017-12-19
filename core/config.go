package core

import (
	io "io/ioutil"
	"github.com/tidwall/gjson"
	"time"
	"github.com/astaxie/beego/logs"
)

var rpcConfig RpcConfig
var zkConfig ZKConfig
var logConfig LogConfig

type AcceptorConfig struct {
	Name    string `json:"name"`
	Service string `json:"service"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Version string `json:"version"`
	Weight  string `json:"weight"`
}

type ConnectorConfig struct {
	Name        string `json:"name"`
	Service     string `json:"service"`
	Version     string `json:"version"`
	Capacity    int32 `json:"capacity"`
	MaxCap      int32 `json:"maxCap"`
	IdleTimeout time.Duration `json:"idleTimeout"`
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

type LogConfig struct {
	FileName string `json:"filename"`
	Level    string `json:"level"`
	Maxlines string `json:"maxlines"`
	Maxsize  string `json:"maxsize"`
	Daily    bool   `json:"daily"`
	Maxdays  string `json:"maxdays"`
}


func Load(filename string) []byte {
	data, err := io.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	datajson := []byte(data)
	return datajson
}

func InitConfig(configPath string) {
	data := Load(configPath + "/rpcServiceConfig.json")
	rpcConfig.Acceptor = AcceptorConfig{}
	rpcConfig.Connector = []ConnectorConfig{}
	gjson.Unmarshal(data, &rpcConfig)

	zkConfig.Zk_option = ZKConfigOption{}
	gjson.Unmarshal(Load(configPath+"/config.json"), &zkConfig)
}

func InitLog(configPath string)  {
	logs.SetLogger("console")
	logData := Load(configPath+"/log.json")
	gjson.Unmarshal(logData, &logConfig)
	logs.SetLogger(logs.AdapterFile, string(logData))
}

func (config *RpcConfig) getRpcConfig() RpcConfig {
	return rpcConfig
}

func (config *ZKConfig) getZkConfig() ZKConfig {
	return zkConfig
}