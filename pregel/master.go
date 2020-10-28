package pregel

import (
	"sync"
)

type Master struct {

	MasterAddr	string
	Nodes 	[]string
	NodesNum	int
	ActiveNum	int
}

func (ma *Master) Schedule() error {
	
	index := 0
	var wg sync.WaitGroup
	for node := range ma.Nodes {
		index = index + 1
		wg.Add(1)
		go func(addr string, index int) {
			defer wg.Done()
			wk := Worker{index, 1, true, addr, ma.MasterAddr, []string {}, []int{1}}
			wk.SuperStep(0)
		}(ma.Nodes[node], index)
	} 
	wg.Wait()
	return nil
}
