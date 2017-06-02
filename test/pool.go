package main

import (
	"net"
	"github.com/fatih/pool"
	"fmt"
)

func main() {
	// create a factory() to be used with channel based pool
	factory := func() (net.Conn, error) { return net.Dial("tcp", "127.0.0.1:3000") }

	// create a new channel based pool with an initial capacity of 5 and maximum
	// capacity of 30. The factory will create 5 initial connections and put it
	// into the pool.
	p, err := pool.NewChannelPool(5, 30, factory)
fmt.Println(err)
	// now you can get a connection from the pool, if there is no connection
	// available it will create a new one via the factory function.
	conn, _ := p.Get()
	conn2, _ := p.Get()
	conn3, _ := p.Get()

	conn4, _ := p.Get()

	conn5, _ := p.Get()

	conn6, _ := p.Get()
	// do something with conn and put it back to the pool by closing the connection
	// (this doesn't close the underlying connection instead it's putting it back
	// to the pool).

	current1 := p.Len()
	fmt.Println(current1)

	conn.Close()
	conn2.Close()
	conn4.Close()
	conn5.Close()
	conn6.Close()

	current2 := p.Len()
	fmt.Println(current2)
	conn3.Close()
	// close pool any time you want, this closes all the connections inside a pool
	p.Close()

	// currently available connections in the pool
	current := p.Len()
	fmt.Println(current)
}
