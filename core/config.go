package core

import (
	io "io/ioutil"
	"github.com/tidwall/gjson"
	"github.com/wangbinlml/go-thrift-rpc/core/logs"
	"time"
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
	// The opened file
	Filename   string `json:"filename"`

	// Rotate at line
	MaxLines         int `json:"maxlines"`

	// Rotate at size
	MaxSize        int `json:"maxsize"`

	// Rotate daily
	Daily         bool  `json:"daily"`
	MaxDays       int64 `json:"maxdays"`
	Rotate bool `json:"rotate"`
	Level int `json:"level"`
	Perm string `json:"perm"`
	Category string `json:"category"`
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
	data := Load(configPath+"/log4go.json")
	gjson.Unmarshal(data, &logConfig)

	config := string(data)
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, config)
	logs.SetCategory(logConfig.Category)

	//日志默认不输出调用的文件名和文件行号
	logs.EnableFuncCallDepth(true)
}

func (config *RpcConfig) getRpcConfig() RpcConfig {
	return rpcConfig
}

func (config *ZKConfig) getZkConfig() ZKConfig {
	return zkConfig
}