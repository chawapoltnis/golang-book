package main

import "fmt"

func main() {
	fmt.Println(add(1,2,3))

	xs :=[]int{1,2,3,5}
	fmt.Println(add(xs...))
	show(xs...)
}

func add(args ...int) int {
	total :=0
	for _,v:=range args {
		total +=v
	}
	return total
}

func show(args ...int) {
	for index,value :=range args {
		fmt.Printf("index :%v , value:%v\n",index,value)
	}
}