package rpc

import (
	"testing"
	"time"
)

var (
	address = "0.0.0.0:9889"
)

func TestRpc_Start(t *testing.T) {
	rpc := NewRpc(nil,nil,nil)
	err := rpc.Start(address)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second*20)
	rpc.Stop()
}