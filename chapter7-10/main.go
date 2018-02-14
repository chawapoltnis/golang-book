package main

import (
	"fmt"
)

func main() {
	maps := map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}

	for i,number := range maps {
		fmt.Println(i,number)
		value, ok :=maps[i]
		fmt.Printf("value:%v and ok:%v\n",value,ok)
	}
}