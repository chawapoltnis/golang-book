package main

import (
	"fmt"
)

func main() {
	//c :=make(chan int)	//memory leak
	c :=make(chan int, 2)	//FIx by use buffered channel
	go func() {c <-1}()
	go func() {c <-2}()
	go func() {c <-3}()
	fmt.Println(<-c)
	fmt.Println(<-c)
}