package core

import (
	"github.com/wangbinlml/go-thrift-rpc/core/gen-go/rpc"
	log "github.com/alecthomas/log4go"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"github.com/youtube/vitess/go/pools"
	"errors"
	"strings"
	"golang.org/x/net/context"
)

var serviceMap = make(map[string]*pools.ResourcePool)
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
						connector.createServer(connectorObj,serviceName, path[0], path[1])
					}
				}
				fmt.Printf("%+v\n", snapshot)
			case err := <-errors:
				fmt.Printf("%+v\n", err)
			}
		}
	}()
}

type ResourceConn struct {
	client *rpc.RPCInvokeServiceClient
}

func (r ResourceConn) Close() {
	log.Info("Transport Close")
	r.client.Transport.Close()
}

func (connector *ThriftConnector) createServer(connectorObj ConnectorConfig, serviceName string, ip string, port string) {

	thriftPool := pools.NewResourcePool(func() (pools.Resource, error) {
		transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		addr := net.JoinHostPort(ip, port)
		transport, err := thrift.NewTSocket(addr) // client端不设置超时
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
			return nil, err
		}

		useTransport := transportFactory.GetTransport(transport)
		client := rpc.NewRPCInvokeServiceClientFactory(useTransport, protocolFactory)
		if err = client.Transport.Open(); err != nil {
			fmt.Fprintln(os.Stderr, "Error opening socket to "+ip+":"+port, " ", err)
			os.Exit(1)
			return nil, err
		}
		return &ResourceConn{client}, err
	}, connectorObj.Capacity, connectorObj.MaxCap, connectorObj.IdleTimeout)
	serviceMap[serviceName] = thriftPool
}

func (connector *ThriftConnector) Invoke(service string, method string, msg *rpc.Msg) (*rpc.Msg, error) {
	if serviceName, ok := sMap[service]; ok {
		if pool, ok := serviceMap[serviceName]; ok {
			//defer pool.Close()
			ctx := context.TODO()
			var rmsg *rpc.Msg
			var resource, el = pool.Get(ctx)
			defer pool.Put(resource)
			if el == nil {
				conn:= resource.(*ResourceConn)
				rmsg, el = conn.client.Invoke(service, method, msg)
			}
			return rmsg, el
		} else {
			fmt.Println("Key Not Found")
			return nil, errors.New("Key Not Found")
		}
	} else {
		fmt.Println("Key Not Found")
		return nil, errors.New("Connector Not Found")
	}
}
