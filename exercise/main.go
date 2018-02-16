package main

import (
	"fmt"
)

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))


	fmt.Println("    _")
	fmt.Println("  |  |")
	fmt.Println("  |  | c")

}

func weatherCelsius(celsius int, desc string) string {
	
	var ret string

	switch celsius {
	case 25: ret=" _  _"+"\n"+" _||_"+"\n"+"|_  _| c"+"\n"+desc
	case 34: ret=" _"+"\n"+" _||_|"+"\n"+" _|  | c"+"\n"+desc
	case 17: ret="    _"+"\n"+"  |  |"+"\n"+"  |  | c"+"\n"+desc
	//case 9: ret=
	default : ret="x"
	}
	return ret
}