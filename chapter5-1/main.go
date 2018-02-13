package main

import "fmt"

func main() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")

	fmt.Println("====use for loop====")
	for i := 1; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println("====for 2====")
	i := 1
	for {
		fmt.Println(i)
		i++
		if i > 3 {
			break
		}
	}

	fmt.Println("====for 3====")
	i=1
	for i<4 {
		fmt.Println(i)
		i++
	}
}
