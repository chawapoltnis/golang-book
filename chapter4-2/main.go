package main

import "fmt"

func main() {
	fmt.Print("Enter a Fahrenheit: ")
	var F float64
	fmt.Scanf("%f", &F)
	//output := input *2
	C := (F - 32) * 5/9
	fmt.Printf("Celcious: %.2f\n",C)	//round float number to 2decimal
}