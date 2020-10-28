package pregel

import (
	"fmt"
	"net/rpc"
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
	for v := range wk.Msg {
		if v>wk.Value {
			wk.Value = v
			wk.Active = true
		}
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