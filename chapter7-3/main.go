package main

import "fmt"

func main() {
	arr :=[5]int{1,2,3,4,5}
	fmt.Println(arr)

	slice := arr[3:5]
	fmt.Println(slice)
	fmt.Printf("length %v, capacity %v, %v\n",len(slice),cap(slice),slice)
}