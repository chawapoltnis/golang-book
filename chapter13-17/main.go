package main

import (
	"time"
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x:=0;x<10;x++{
			naturals <-x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
		//for {
			//x :=<-naturals
			squares <-x*x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
		time.Sleep(1*time.Second)
	}
}