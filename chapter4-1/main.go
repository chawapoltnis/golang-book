package main

import "fmt"

var j string="global variable"

func main() {
	//Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	//Short Declaration
	//Type Inference
	z := "Hello, World"
	fmt.Print(z)
	fmt.Printf(" Type:%T\n", z)

	n := 3987
	fmt.Print(n)
	fmt.Printf(" Type:%T\n", n)

	m := 3987.5
	fmt.Print(m)
	fmt.Printf(" Type:%T\n", m)

	f()

	//cannot assign to constance variable
	const k string = "Hello, k"
	//k = "Other string"

	//define multiple variable
	var (
		aa=5
		ab=6
		ac=7
	)
	fmt.Println(aa,ab,ac)

	//define multiple variable short term
	a1,a2 := 5,"heck"
	fmt.Print(a1,",",a2)
	fmt.Printf(" Type:%T,%T\n",a1,a2)

	//swaping variable
	bL,bR := "Left","Right"
	bL,bR = bR,bL
	fmt.Println(bL,bR)
}

func f() {
	fmt.Println(j)
}