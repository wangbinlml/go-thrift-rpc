package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	log "github.com/alecthomas/log4go"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"strings"
	"errors"
)

var serviceMap = make(map[string]*rpc.RPCInvokeServiceClient)
var sMap = make(map[string]string)

type ThriftConnector struct {
	Config []ConnectorConfig
}

func (connector *ThriftConnector) init(c []ConnectorConfig) {
	connector.Config = c
	log.Info("ThriftConnector init")
}

func (connector *ThriftConnector) start() {
	connector.createConnect()
	log.Info("ThriftConnector start")
}

func (connector *ThriftConnector) createConnect() {
	for _, v := range connector.Config {
		configObj := ConnectorConfig(v)
		sMap[configObj.Name] = configObj.Service
		connector.initThrift(configObj)
	}
}
func (connector *ThriftConnector) initThrift(connectorObj ConnectorConfig) {
	var service = connectorObj.Service
	var version = connectorObj.Version
	if version == "" {
		version = "1.0.0"
	}
	var path = service + "/" + version
	snapshots, errors := zkUtil.GetNodesW(path)
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				if len(snapshot) > 0 {
					serviceName := connectorObj.Service
					for _, url := range snapshot {
						path := strings.Split(url, ":")
						log.Info("connector path " + path[0] + ":" + path[1])
						connector.createServer(serviceName, path[0], path[1])
					}
				}
				fmt.Printf("%+v\n", snapshot)
			case err := <-errors:
				fmt.Printf("%+v\n", err)
			}
		}
	}()
}

func (connector *ThriftConnector) createServer(serviceName string, ip string, port string) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(ip, port))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRPCInvokeServiceClientFactory(useTransport, protocolFactory)
	serviceMap[serviceName] = client
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+ip+":"+port, " ", err)
		os.Exit(1)
	}
	defer transport.Close()
}

func (connector *ThriftConnector) Invoke(service string, method string, msg *rpc.Msg) (*rpc.Msg, error) {
	if serviceName, ok := sMap[service]; ok {
		if client, ok := serviceMap[serviceName]; ok {
			return client.Invoke(service, method, msg)
		} else {
			fmt.Println("Key Not Found")
			return nil, errors.New("Key Not Found")
		}
	} else {
		fmt.Println("Key Not Found")
		return nil, errors.New("Connector Not Found")
	}
}
