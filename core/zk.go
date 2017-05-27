package core

import (
	log "github.com/alecthomas/log4go"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
	"errors"
	"path"
	"syscall"
	"os"
)

var (
	// error
	ErrNoChild      = errors.New("zk: children is nil")
	ErrNodeNotExist = errors.New("zk: node not exist")
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Connect connect to zookeeper, and start a goroutine log the event.
func Connect(addr []string, timeout time.Duration) (*zk.Conn, error) {
	conn, _, err := zk.Connect(addr, timeout)
	if err != nil {
		log.Error("zk.Connect(\"%v\", %d) error(%v)", addr, timeout, err)
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
func Create(conn *zk.Conn, fpath string, data string, flags int32) error {
	// create zk root path
	tpath := ""
	for _, str := range strings.Split(fpath, "/")[1:] {
		tpath = path.Join(tpath, "/", str)
		log.Debug("create zookeeper path: \"%s\"", tpath)
		acl := zk.WorldACL(zk.PermAll)
		_, err := conn.Create(tpath, []byte(data), flags, acl)
		if err != nil {
			if err == zk.ErrNodeExists {
				log.Warn("zk.create(\"%s\") exists", tpath)
			} else {
				log.Error("zk.create(\"%s\") error(%v)", tpath, err)
				return err
			}
		}
	}

	return nil
}
func CreateProtectedEphemeralSequential(conn *zk.Conn, fpath string, data string) (string, error) {
	acl := zk.WorldACL(zk.PermAll)
	return conn.CreateProtectedEphemeralSequential(fpath, []byte(data), acl)
}

// RegisterTmp create a ephemeral node, and watch it, if node droped then send a SIGQUIT to self.
func RegisterTemp(conn *zk.Conn, fpath string, data []byte) error {
	tpath, err := conn.Create(path.Join(fpath)+"/", data, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Error("conn.Create(\"%s\", \"%s\", zk.FlagEphemeral|zk.FlagSequence) error(%v)", fpath, string(data), err)
		return err
	}
	log.Debug("create a zookeeper node:%s", tpath)
	// watch self
	go func() {
		for {
			log.Info("zk path: \"%s\" set a watch", tpath)
			exist, _, watch, err := conn.ExistsW(tpath)
			if err != nil {
				log.Error("zk.ExistsW(\"%s\") error(%v)", tpath, err)
				log.Warn("zk path: \"%s\" set watch failed, kill itself", tpath)
				killSelf()
				return
			}
			if !exist {
				log.Warn("zk path: \"%s\" not exist, kill itself", tpath)
				killSelf()
				return
			}
			event := <-watch
			log.Info("zk path: \"%s\" receive a event %v", tpath, event)
		}
	}()
	return nil
}

// GetNodesW get all child from zk path with a watch.
func GetNodesW(conn *zk.Conn, path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)

	go func() {
		for {
			snapshot, _, events, err := conn.ChildrenW(path)
			if err != nil {
				log.Error("zk.ChildrenW(\"%s\") error(%v)", path, err)
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
func GetNodes(conn *zk.Conn, path string) ([]string, error) {
	nodes, stat, err := conn.Children(path)
	if err != nil {
		if err == zk.ErrNoNode {
			return nil, ErrNodeNotExist
		}
		log.Error("zk.Children(\"%s\") error(%v)", path, err)
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

func Delete(conn *zk.Conn, path string) error {
	return conn.Delete(path, 0)
}
func Exists(conn *zk.Conn, path string) (bool, error) {
	b, _, e := conn.Exists(path)
	return b, e;
}

func killSelf() {
	if err := syscall.Kill(os.Getpid(), syscall.SIGQUIT); err != nil {
		log.Error("syscall.Kill(%d, SIGQUIT) error(%v)", os.Getpid(), err)
	}
}
