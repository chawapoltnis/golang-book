package main

import (
	"fmt"
)

func main() {
	addFunc := func(a int) (func(b int)int) {
		return func(b int) int {
			return a+b
		}
	}

	add2With := addFunc(2)
	fmt.Println(add2With(3))
}