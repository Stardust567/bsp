package pregel

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
)

func (wk *Worker) ValueRecv(value int, reply *int) error {
	fmt.Println("Worker", wk.Number, "ValueRecv: ", value)
	wk.Msg = append(wk.Msg, value)
	return nil
}

func (wk *Worker) StartWorker() error {
	rpc.Register(new(Worker))
	// ===== workaround ==========
    oldMux := http.DefaultServeMux
    mux := http.NewServeMux()
    http.DefaultServeMux = mux
	// ===========================
	
	rpc.HandleHTTP()

	// ===== workaround ==========
    http.DefaultServeMux = oldMux
    // ===========================
	listener, err := net.Listen("tcp", wk.Myaddr)
	if err != nil {
		fmt.Println("listen error:", err)
	}
	fmt.Println("Worker", wk.Number, " is listening in ", wk.Myaddr)
	go http.Serve(listener, nil)
	return nil
}

func (wk *Worker) Call(addr string) error {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		fmt.Println("dialing:", err)
	}
	fmt.Println("Worker", wk.Number, "Call", addr)
	var reply int
	err = client.Call("Worker.ValueRecv", wk.Value, &reply)
	if err != nil {
		fmt.Println("rpc call error:", err)
	}
	return nil
}