package main

import (
	"fmt"
	"time"
)

func main() {
	volumn :=100
	start := time.Now()

	container := order(volumn)
	for _,cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

/*
func order(volumn int) (container []string) {
	for i:=1; i<=volumn; i++ {
		//cashier receive order
		time.Sleep(5*time.Microsecond)
		coffee:= fmt.Sprintf("order: %d", i)

		//barista brew coffee
		time.Sleep(100*time.Millisecond)
		coffee = fmt.Sprintf("%s %s",coffee, "espresso")

		//waiter serve coffee
		time.Sleep(5 * time.Millisecond)
		container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))
	}
	return
}*/

func order(volumn int) (container []string) {

	cashier := make(chan string)
	barista1 := make(chan string)
	barista2 := make(chan string)
	waiter := make(chan string)

		//cashier receive order
	go docashier(volumn,cashier)
	go dobarista1(waiter,cashier,barista1)
	go dobarista2(waiter,cashier,barista2)
	//go dowaiter(waiter,barista1,barista2)

	for x := range waiter {
		container = append(container, x)
	}

	return
}

func docashier(vol int,in chan<- string) {
	for x:=1;x<=vol;x++{
		time.Sleep(5*time.Microsecond)
		in<- fmt.Sprintf("order: %d", x)
	}
	close(in)
}

func dobarista1(out2 chan<- string,in <-chan string, out chan string) {
	for x := range in {
		time.Sleep(100*time.Millisecond)
		out<- fmt.Sprintf("%s %s",x, "espresso")
		dowaiter(out2,<-out)
	}
	close(out)
}

func dobarista2(out2 chan<- string,in <-chan string, out chan string) {
	for x := range in {
		time.Sleep(100*time.Millisecond)
		out<- fmt.Sprintf("%s %s",x, "espresso")
		dowaiter(out2,<-out)
	}
	close(out)
}

func dowaiter(out chan<- string,coffee string) {
	time.Sleep(5*time.Millisecond)
	out <- fmt.Sprintf("%s %s", coffee, "ready :)")
}
