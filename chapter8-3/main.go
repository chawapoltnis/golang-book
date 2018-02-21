package main

import (
	"fmt"
)

func main() {
	//array := [3]int{1,2,3}
	//double(array)
	//fmt.Printf("original addr %p\n", &array)
	//fmt.Printf("original %v\n", array)

	array2 := [3]int{1,2,3}
	double2(&array2)
	fmt.Printf("original2 addr %p\n", &array2)
	fmt.Printf("original2 %v\n", array2)
}

func double(nums [3]int) {
	fmt.Printf("double addr %p\n", &nums)
	for i:= range nums {
		nums[i] *=2
	}
	fmt.Println(nums)
}

func double2(nums *[3]int) {
	fmt.Printf("double2 addr %p\n", nums)
	for i:= range nums {
		nums[i] *=2
	}
	fmt.Println(*nums)
}