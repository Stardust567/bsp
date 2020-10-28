package pregel

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
)

func (ma *Master) InactiveRecv(value int, reply *int) error {
	ma.ActiveNum += 1
	fmt.Println(value)
	return nil
}

func (ma *Master) StartServer() error {
	mr := new(Master)
	rpc.Register(mr)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ma.MasterAddr)
	if err != nil {
		fmt.Println("listen error:", err)
	}
	go http.Serve(listener, nil)
	return nil
}