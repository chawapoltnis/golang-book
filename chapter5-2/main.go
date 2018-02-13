package main

import "fmt"

func main() {
	for number := 1; number <= 18; number++ {

		switch {
		case number%15 == 0:
			fmt.Println("FizzBuzz")
		case number%5 == 0:
			fmt.Println("Buzz")
		case number%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(number)
		}

		switch number%2 == 0 {
		case true:
			fmt.Println("even")
		default:
			fmt.Println("odd")
		}
	}
}
