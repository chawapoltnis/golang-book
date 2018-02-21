package main

import (
	"fmt"
	"strconv"
)

func main() {
	for number := 1; number <= 20; number++ {

		fmt.Println(number,fizzbuzz(number))

	}
}

func fizzbuzz(n int) string {

	ln :=[3]int{15,3,5}
	str :=[3]string{"FizzBuzz","Fizz","Buzz"}
	for i:=0;i<len(ln);i++ {
		if n%ln[i]==0{
			return str[i]
		}
	}
	return strconv.Itoa(n)
}