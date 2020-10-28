package main

import (
    "bsp/pregel"
)

func pregel_func() {
    ma := pregel.Master{"127.0.0.1:1234", []string {"127.0.0.1:8890", "127.0.0.1:8891"}, 1, 0}
    ma.StartServer()
    ma.Schedule()
}

func main() {
    pregel_func()
}