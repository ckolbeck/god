package common

import (
	"bytes"
	"fmt"
	"net/rpc"
)

type Remotes []Remote
func (self Remotes) Equal(other []Remote) bool {
	if len(self) != len(other) {
		return false
	}
	for i, r := range self {
		if !r.Equal(other[i]) {
			return false
		}
	}
	return true
}

type Remote struct {
	Pos  []byte
	Addr string
}

func (self Remote) Equal(other Remote) bool {
	return self.Addr == other.Addr && bytes.Compare(self.Pos, other.Pos) == 0
}
func (self Remote) Less(other Remote) bool {
	val := bytes.Compare(self.Pos, other.Pos)
	if val == 0 {
		val = bytes.Compare([]byte(self.Addr), []byte(other.Addr))
	}
	return val < 0
}
func (self Remote) String() string {
	return fmt.Sprintf("[%v@%v]", HexEncode(self.Pos), self.Addr)
}
func (self Remote) Call(service string, args, reply interface{}) error {
	return Switch.Call(self.Addr, service, args, reply)
}
func (self Remote) Go(service string, args, reply interface{}) *rpc.Call {
	return Switch.Go(self.Addr, service, args, reply)
}