package core

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *zk.Conn {
	servers := strings.Split("127.0.0.1:2181", ",")
	conn, _, err := zk.Connect(servers, time.Second)
	must(err)
	return conn
}

func mirror(conn *zk.Conn, path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)

	go func() {
		for {
			snapshot, _, events, err := conn.ChildrenW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- snapshot
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
		}
	}()

	return snapshots, errors
}

func main() {
	conn1 := connect()
	defer conn1.Close()

	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	_, err := conn1.Create("/mirror", nil, int32(0), acl)
	if err != nil {
		fmt.Printf("create: %+v\n", err)
	}

	snapshots, errors := mirror(conn1, "/mirror")
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				fmt.Printf("%+v\n", snapshot)
			case err := <-errors:
				fmt.Printf("%+v\n", err)
			}
		}
	}()

	conn2 := connect()
	time.Sleep(time.Second)

	_, err = conn2.Create("/mirror/one", []byte("one"), flags, acl)
	must(err)
	time.Sleep(time.Second)

	_, err = conn2.Create("/mirror/two", []byte("two"), flags, acl)
	must(err)
	time.Sleep(time.Second)

	err = conn1.Delete("/mirror/two", 0)
	must(err)
	time.Sleep(time.Second)

	_, err = conn2.Create("/mirror/three", []byte("three"), flags, acl)
	must(err)
	time.Sleep(time.Second)

	conn2.Close()
	time.Sleep(time.Second)
}
