package main

import "fmt"

func main() {
	fmt.Println(f1())
	fmt.Println(f3())
}

func f1() int {
	return f2()
}

func f2() int {				//chapter6-1 return from function
	return 1
}

func f3() (r int) {			//chapter6-2 return types can have name
	r=2
	return
}
