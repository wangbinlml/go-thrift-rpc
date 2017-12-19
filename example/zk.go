package main

import (
	"fmt"
	_ "github.com/wangbinlml/go-thrift-rpc/core"
	"time"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/wangbinlml/go-thrift-rpc/core"
	"github.com/astaxie/beego/logs"
)

func main() {

	var zkUtil = new(core.ZKUtil)

	flags := int32(zk.FlagEphemeral)

	err := zkUtil.Create("/mirror", "", int32(0))
	if err != nil {
		fmt.Printf("create: %+v\n", err)
	}

	snapshots, errors := zkUtil.GetNodesW("/mirror")
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				logs.Info("sssssssssssssss")
				fmt.Printf("%+v\n", snapshot)
			case err := <-errors:
				fmt.Printf("%+v\n", err)
			}
		}
	}()

	time.Sleep(time.Second)

	err = zkUtil.Create("/mirror/one", "one", flags)
	core.Must(err)
	time.Sleep(time.Second)

	err = zkUtil.Create("/mirror/two", "two", flags)
	core.Must(err)
	time.Sleep(time.Second)

	err = zkUtil.Delete( "/mirror/two")
	core.Must(err)
	time.Sleep(time.Second)

	err = zkUtil.Create("/mirror/three", "three", flags)
	core.Must(err)
	time.Sleep(time.Second)

	time.Sleep(time.Second)
}
