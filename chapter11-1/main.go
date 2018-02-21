package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	w float64
	h float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

type measure interface {
	area() float64
}

func main() {
	c := &Circle{0,0,5}		//not same as call Circle
	printArea(c)

	r := Rectangle{3,4}		//same as call &Rectangle
	printArea(r)
}

func (r Rectangle) area() float64{
	return r.w * r.h
}

func (c *Circle) area() float64{
	return math.Pi *c.r *c.r
}

func printArea(m measure) {
	fmt.Println(m.area())
}

