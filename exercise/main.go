package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak many"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
	fmt.Println(weather2(1234567890, "nhor"))

}

func weatherCelsius(celsius int, desc string) string {

	var ret string

	switch celsius {
	case 25:
		ret = " _  _" + "\n" + " _||_ " + "\n" + "|_  _| c" + "\n" + desc
	case 34:
		ret = " _   " + "\n" + " _||_|" + "\n" + " _|  | c" + "\n" + desc
	case 17:
		ret = "    _" + "\n" + "  |  |" + "\n" + "  |  | c" + "\n" + desc
	//case 9: ret=
	default:
		ret = "x"
	}
	return ret
}

func weather2(celsius int, desc string) string {
	var ret string
	cel := strconv.Itoa(celsius)

	num1 := []string{"   ", "  |", "  |"}
	num2 := []string{" _ ", " _|", "|_ "}
	num3 := []string{" _ ", " _|", " _|"}
	num4 := []string{"   ", "|_|", "  |"}
	num5 := []string{" _ ", "|_ ", " _|"}
	num6 := []string{" _ ", "|_ ", "|_|"}
	num7 := []string{" _ ", "  |", "  |"}
	num8 := []string{" _ ", "|_|", "|_|"}
	num9 := []string{" _ ", "|_|", " _|"}
	num0 := []string{" _ ", "| |", "|_|"}

	var str [][]string
	for _,ascii := range cel {
		numi:=string(ascii)
		switch numi {
		case "1":str=append(str,num1)
		case "2":str=append(str,num2)
		case "3":str=append(str,num3)
		case "4":str=append(str,num4)
		case "5":str=append(str,num5)
		case "6":str=append(str,num6)
		case "7":str=append(str,num7)
		case "8":str=append(str,num8)
		case "9":str=append(str,num9)
		default :str=append(str,num0)
	    }
	}

	for n:=0;n<3;n++ {
		for m:=0;m<len(cel);m++ {
			ret+=str[m][n]
		}
		ret+="\n"
	}
	ret+=desc

	return ret
}
