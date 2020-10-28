package pregel

import (
	//"fmt"
	"sync"
	"time"
)

type Master struct {

	MasterAddr	string
	NodesAddr 	[]string
	ActiveNum	int

}

func (ma *Master) Schedule(file map[int][]string) error {
	
	index := 0
	var wg sync.WaitGroup
	for node := range ma.NodesAddr {
		index = index + 1
		wg.Add(1)
		go func(addr string, index int) {
			defer wg.Done()
			wk := Worker{index, 1+index, true, addr, ma.MasterAddr, file[index], []int{1}}
			wk.Barrier()
			time.Sleep(10000)
			wk.SuperStep(0)
		}(ma.NodesAddr[node], index)
	} 
	wg.Wait()
	return nil
}
