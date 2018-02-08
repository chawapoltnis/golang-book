package main

import "fmt"

func main() {
	fmt.Print("Enter a Feet: ")
	var F float64
	fmt.Scanf("%f", &F)
	M := F*0.3048
	fmt.Printf("Meters: %.4f\n",M)
}