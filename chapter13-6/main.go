package main

import (
	"sync/atomic"
	"fmt"
	"sync"
)

var (
	counter int64
	wg sync.WaitGroup
)
func main() {
	wg.Add(16)
	for i:=1;i<=16;i++ {
		go increment(i)
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
}

func increment(n int){
	defer wg.Done()
	for count :=0; count<2 ; count++{
		atomic.AddInt64(&counter,1)
	}
}