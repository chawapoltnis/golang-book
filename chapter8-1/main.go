package main

import (
	"fmt"
)

func main() {
	amount :=5
	double(amount)
	fmt.Printf("original %v\n",amount)

	amount2 :=6
	double2(&amount2)
	fmt.Printf("original2 %v\n",amount2)
}

func double(number int) {
	number *=2
	fmt.Println(number)
}

func double2(number *int) {
	*number *=2
	fmt.Println(*number)
}