package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	 var A Circle
	 fmt.Printf("Area A %v\n",A.area())
	 A.changeRedius(2)
	 fmt.Printf("Area A update: %v\n",A.area())

	 littleC := Circle{0,0,5}
	 fmt.Println("littleC ",littleC.area())
	 littleC.changeRedius(10)
	 fmt.Println("littleC ",littleC.area())

	 bigC := &Circle{0,0,5}
	 fmt.Println("bicC ",bigC.area())
	 bigC.changeRedius(10)
	 fmt.Println("bigC ",bigC.area())
}

func (c Circle) area() float64{
	return math.Pi * c.r * c.r
}

func (c *Circle) changeRedius(r float64) {
	c.r=r
}