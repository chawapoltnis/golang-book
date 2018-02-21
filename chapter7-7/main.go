package main

import (
	"fmt"
)

func main() {
	mymap:=make(map[int]int)
	mymap[1]=1
	mymap[2]=2
	//mymap[3]=9

	fmt.Println(mymap[3])
	if mymap[3]!=0 {
		fmt.Println(mymap[3])
	}

	//map has return ok flag
	if value, ok := mymap[3]; ok {
		fmt.Println(value)
	}

	value, ok :=mymap[3]
	fmt.Printf("type of ok is:%T\n",ok)
	if ok {
		fmt.Println(value)
	}
}