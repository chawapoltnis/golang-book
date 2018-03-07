package main

import (
	"fmt"
	"time"
)

func main() {
	cl := make(chan string, 1)
	go func() {
		time.Sleep(500*time.Millisecond)
		cl <- "result 1"
	}()

	select {
	case res := <-cl:
		fmt.Println(res)
	case <-time.After(5*time.Second):
		fmt.Println("timeout 1")
	}
}