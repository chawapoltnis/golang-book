package main

import (
	"fmt"
)

func main() {
	var addVar func(int,int) int
	addVar = func(a,b int) int {
		return a+b
	}
	fmt.Println(addVar(2,3))

	func(a,b int) int {
		fmt.Println(a+b)
		return a+b
	}(3,4)

	fmt.Println(
		func(a,b int)int{
			return a+b
		}(4,5))
}