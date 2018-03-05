package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i:=0;i<4; i++ {
		go func(n int){
			defer wg.Done()		//do at last
			for i:=0;i<10;i++{
				fmt.Println(n,":",i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished")
}
