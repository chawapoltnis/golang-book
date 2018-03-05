package main

import (
	"fmt"
	"strconv"
)

func main() {
	//for number := 1; number <= 20; number++ {
	//	fmt.Println(number,func(n int)string {
	//		ln :=[3]int{15,3,5}
	//		str :=[3]string{"FizzBuzz","Fizz","Buzz"}
	//		for i:=0;i<len(ln);i++ {
	//			if n%ln[i]==0{
	//				return str[i]
	//			}
	//		}
	//		return strconv.Itoa(n)
	//	}(number))
	//}

	for number := 1; number <= 20; number++ {
		fmt.Println(number,fizzbuzz(number))
	}
}

func fbTemplate(fbnumber int, str string) (func(int) (string, bool)){
	return func(n int)(string,bool){
		if n%fbnumber ==0 {
			return str, true
		}
		return "",false
	}
}

func fizzbuzz(number int) string {
	fbArray :=[...]func(n int) (string ,bool){
		fbTemplate(15, "FizzBuzz"),
		fbTemplate(5, "Buzz"),
		fbTemplate(3, "Fizz"),
	}

	for i:=0;i<len(fbArray);i++{
		if str, ok :=fbArray[i](number); ok {
			return str
		}
	}
	return strconv.Itoa(number)
}
	

