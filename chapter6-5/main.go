package main

import (
	"fmt"
	"strconv"
)

func main() {
	for number := 1; number <= 18; number++ {

		fmt.Println(number,fizzbuzz(number))

	}
}

func fizzbuzz(n int) (str string) {

	switch {
	case n%15 == 0:
		str="FizzBuzz"
	case n%5 == 0:
		str="Buzz"
	case n%3 == 0:
		str="Fizz"
	default:
		str=strconv.Itoa(n)
	}
	return
}