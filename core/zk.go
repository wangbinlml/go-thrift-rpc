package core

import (
	zookeeper "github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
	"errors"
	"path"
	"syscall"
	"os"
	"github.com/astaxie/beego/logs"
)

var (
	// error
	ErrNoChild      = errors.New("zk: children is nil")
	ErrNodeNotExist = errors.New("zk: node not exist")
)

type ZKUtil struct {
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

var conn *zookeeper.Conn

func InitZk() {
	v := new(ZKConfig)
	var zkConfig = v.getZkConfig()
	servers := strings.Split(zkConfig.Zk_path, ",")
	var zkUtil = new(ZKUtil)
	conn, _ = zkUtil.Connect(servers, time.Second)
}

// Connect connect to zookeeper, and start a goroutine log the event.
func (zkUtil *ZKUtil) Connect(addr []string, timeout time.Duration) (*zookeeper.Conn, error) {
	conn, _, err := zookeeper.Connect(addr, timeout)
	if err != nil {
		logs.Error("zk.Connect(\"%v\", %d) error(%v)", addr, timeout, err)
		return nil, err
	}
	/*go func() {
		for {
			event := <-session
			log.Debug("zookeeper get a event: %s", event.State.String())
		}
	}()*/
	return conn, nil
}

// Create create zookeeper path, if path exists ignore error
func (zkUtil *ZKUtil) Create(fpath string, data string, flags int32) error {
	// create zk root path
	tpath := ""
	for _, str := range strings.Split(fpath, "/")[1:] {
		tpath = path.Join(tpath, "/", str)
		logs.Debug("create zookeeper path: \"%s\"", tpath)
		acl := zookeeper.WorldACL(zookeeper.PermAll)
		_, err := conn.Create(tpath, []byte(data), flags, acl)
		if err != nil {
			if err == zookeeper.ErrNodeExists {
				logs.Warn("zk.create(\"%s\") exists", tpath)
			} else {
				logs.Error("zk.create(\"%s\") error(%v)", tpath, err)
				return err
			}
		}
	}

	return nil
}
func (zkUtil *ZKUtil) CreateProtectedEphemeralSequential(fpath string, data string) (string, error) {
	acl := zookeeper.WorldACL(zookeeper.PermAll)
	return conn.CreateProtectedEphemeralSequential(fpath, []byte(data), acl)
}

// RegisterTmp create a ephemeral node, and watch it, if node droped then send a SIGQUIT to self.
func (zkUtil *ZKUtil) RegisterTemp(fpath string, data []byte) error {
	tpath, err := conn.Create(path.Join(fpath)+"/", data, zookeeper.FlagEphemeral|zookeeper.FlagSequence, zookeeper.WorldACL(zookeeper.PermAll))
	if err != nil {
		logs.Error("conn.Create(\"%s\", \"%s\", zk.FlagEphemeral|zk.FlagSequence) error(%v)", fpath, string(data), err)
		return err
	}
	logs.Debug("create a zookeeper node:%s", tpath)
	// watch self
	go func() {
		for {
			logs.Info("zk path: \"%s\" set a watch", tpath)
			exist, _, watch, err := conn.ExistsW(tpath)
			if err != nil {
				logs.Error("zk.ExistsW(\"%s\") error(%v)", tpath, err)
				logs.Warn("zk path: \"%s\" set watch failed, kill itself", tpath)
				killSelf()
				return
			}
			if !exist {
				logs.Warn("zk path: \"%s\" not exist, kill itself", tpath)
				killSelf()
				return
			}
			event := <-watch
			logs.Info("zk path: \"%s\" receive a event %v", tpath, event)
		}
	}()
	return nil
}

// GetNodesW get all child from zk path with a watch.
func (zkUtil *ZKUtil) GetNodesW(path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)

	go func() {
		for {
			snapshot, _, events, err := conn.ChildrenW(path)
			if err != nil {
				logs.Error("zk.ChildrenW(\"%s\") error(%v)", path, err)
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

// GetNodes get all child from zk path.
func (zkUtil *ZKUtil) GetNodes(path string) ([]string, error) {
	nodes, stat, err := conn.Children(path)
	if err != nil {
		if err == zookeeper.ErrNoNode {
			return nil, ErrNodeNotExist
		}
		logs.Error("zk.Children(\"%s\") error(%v)", path, err)
		return nil, err
	}
	if stat == nil {
		return nil, ErrNodeNotExist
	}
	if len(nodes) == 0 {
		return nil, ErrNoChild
	}
	return nodes, nil
}

func (zkUtil *ZKUtil) Delete(path string) error {
	return conn.Delete(path, 0)
}

func (zkUtil *ZKUtil) Exists(path string) (bool, error) {
	b, _, e := conn.Exists(path)
	return b, e;
}

func (zkUtil *ZKUtil) Close() {
	conn.Close()
	logs.Info("zk client close")
}

func killSelf() {
	if err := syscall.Kill(os.Getpid(), syscall.SIGQUIT); err != nil {
		logs.Error("syscall.Kill(%d, SIGQUIT) error(%v)", os.Getpid(), err)
	}
}
