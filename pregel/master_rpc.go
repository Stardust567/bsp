package pregel

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
)

func (ma *Master) InactiveRecv(value int, reply *int) error {
	ma.ActiveNum += 1
	fmt.Println("Master: InactiveRecv", value)
	return nil
}

func (ma *Master) StartMaster() error {
	rpc.Register(new(Master))
	// ===== workaround ==========
    oldMux := http.DefaultServeMux
    mux := http.NewServeMux()
    http.DefaultServeMux = mux
	// ===========================
	
	rpc.HandleHTTP()
	
	// ===== workaround ==========
    http.DefaultServeMux = oldMux
    // ===========================
	listener, err := net.Listen("tcp", ma.MasterAddr)
	if err != nil {
		fmt.Println("listen error:", err)
	}
	go http.Serve(listener, nil)
	return nil
}