package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	randomnumber := rand.Intn(200)

	var input int
	for count := 1; count <= 5; count++ {
		fmt.Print("Guess what Number: ")
		fmt.Scanf("%d\n", &input)
		//fmt.Println(input)
		/*
			switch {
			case input == randomnumber:
				fmt.Println("correct.")
			case input > randomnumber:
				fmt.Println("too much.")
			default:
				fmt.Println("too small.")
			}*/

		if input == randomnumber {
			fmt.Println("correct")
			break
		}
		if input > randomnumber {
			fmt.Println("too much")
		}
		if input < randomnumber {
			fmt.Println("too small")
		}
	}

	fmt.Printf("correct answer is %d", randomnumber)
}
