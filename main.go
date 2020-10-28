package main

import (
    "bsp/pregel"
)

func pregel_func() {
    ma := pregel.Master{"127.0.0.1:1234", []string {"127.0.0.1:8890", "127.0.0.1:8891"}, 0}
    file := map[int][]string{1:{"127.0.0.1:8891"}, 2:{"127.0.0.1:8890"}}
    ma.StartMaster()
    ma.Schedule(file)
}

func main() {
    pregel_func()
}