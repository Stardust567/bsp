package pregel

import (
	"fmt"
	"net/rpc"
	"sync"
	"time"
)

type Worker struct {

	Number 	int
	Value 	int
	Active	bool
	Myaddr	string
	Master	string 
	Addr 	[]string
	Msg 	[]int
}

func (wk *Worker) SuperStep(SSNumber int) error {
	fmt.Printf("SuperStep %d: Worker %d - value %d\n", SSNumber, wk.Number, wk.Value)
	wk.Active = false
	for i := range wk.Msg {
		if wk.Msg[i]>wk.Value {
			wk.Value = wk.Msg[i]
			wk.Active = true
		}
		wk.Msg = append(wk.Msg[:i], wk.Msg[i+1:]...)
	}
	if wk.Active==false {
		client, err := rpc.DialHTTP("tcp", wk.Master)
		if err != nil {
			fmt.Println("dialing:", err)
		}
		var reply int
		err = client.Call("Master.InactiveRecv", wk.Value, &reply)
		if err != nil {
			fmt.Println("rpc call error:", err)
		}
	}
	return nil
}

func (wk *Worker) Barrier() error {
	fmt.Println("Worker-", wk.Number, "-Barrier")
	wk.StartWorker()
	time.Sleep(10000)
	var wg sync.WaitGroup
	wg.Add(len(wk.Addr))
	for i := range wk.Addr {
		go func(){
			defer wg.Done()
			wk.Call(wk.Addr[i])
		}()
	}
	return nil
}
