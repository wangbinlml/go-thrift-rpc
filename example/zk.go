package main

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"strings"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	servers := strings.Split("127.0.0.1:2181", ",")
	conn1, _ := core.Connect(servers, time.Second)
	defer conn1.Close()

	flags := int32(zk.FlagEphemeral)

	err := core.Create(conn1, "/mirror", "", int32(0))
	if err != nil {
		fmt.Printf("create: %+v\n", err)
	}

	snapshots, errors := core.GetNodesW(conn1, "/mirror")
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				log.Info("sssssssssssssss")
				fmt.Printf("%+v\n", snapshot)
			case err := <-errors:
				fmt.Printf("%+v\n", err)
			}
		}
	}()

	conn2, _ := core.Connect(servers, time.Second)
	defer conn2.Close()
	time.Sleep(time.Second)

	err = core.Create(conn2, "/mirror/one", "one", flags)
	core.Must(err)
	time.Sleep(time.Second)

	err = core.Create(conn2, "/mirror/two", "two", flags)
	core.Must(err)
	time.Sleep(time.Second)

	err = core.Delete(conn2, "/mirror/two")
	core.Must(err)
	time.Sleep(time.Second)

	err = core.Create(conn2, "/mirror/three", "three", flags)
	core.Must(err)
	time.Sleep(time.Second)

	time.Sleep(time.Second)
}
