package main

import (
	"runtime"
	"fmt"
	"sync"
)

var (
	counter int
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
		value := counter
		runtime.Gosched()		//return resource
		value++
		counter=value
	}
}