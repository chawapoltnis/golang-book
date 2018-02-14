package main

import (
	"fmt"
)

func main() {
	for i,ascii := range "golang"{
		fmt.Println(i,ascii)
		fmt.Printf("%v\n",string(ascii))
	}
}