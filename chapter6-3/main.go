package main

import "fmt"

func main() {
	x, y := f()
	fmt.Println(x,y)
	fmt.Println(f2())
}

func f() (int, int) {
	return 3,4
}

func f2() (a int, b int) {
	a=5
	b=6
	//return			//return a,b
	return 7,8			//override a,b
}