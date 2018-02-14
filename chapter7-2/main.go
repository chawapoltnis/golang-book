package main

import "fmt"

func main() {
	slice := make([]int,3,10)
	slice[0]=1
	slice[1]=2
	slice[2]=3
	fmt.Println(slice)

	slice2 :=[]int{1,2,3,4,6}
	fmt.Println(slice2)

	fmt.Println("Slice with length and capacity")
	fmt.Printf("length %v, capacity %v, %v\n",len(slice),cap(slice),slice)

	//append
	for i:=4;i<15;i++{
		slice = append(slice,i)
	}
	fmt.Printf("length %v, capacity %v, %v\n",len(slice),cap(slice),slice)	
}