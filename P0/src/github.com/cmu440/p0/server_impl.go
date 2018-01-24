// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"strconv"
)
type keyValueServer struct {
	// TODO: implement this!
	listener net.Listener
	protocol string
	count int
}

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	// TODO: implement this!
	return keyValueServer{protocol:"tcp",count:0 };
}

func (kvs *keyValueServer) Start(port int) error {
	// TODO: implement this!
	kvs.listener ,err := net.Listen(kvs.protocol, "127.0.0.1:"+strconv.Itoa(port));
	return err
}


func (kvs *keyValueServer) Close() {
	// TODO: implement this!
	if kvs.listener !=nil {
		kvs.listener.Close()
		kvs.listener=nil
	}
	
		
	
}

func (kvs *keyValueServer) Count() int {
	// TODO: implement this!
	
	return kvs.count
}

// TODO: add additional methods/functions below!
