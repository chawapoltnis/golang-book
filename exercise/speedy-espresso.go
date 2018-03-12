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
	barista := make(chan string,2)
	waiter := make(chan string)

		//cashier receive order
	go docashier(volumn,cashier)

	maxbarista :=2
	for i:=1;i<=maxbarista;i++ {
		go dobarista(cashier,barista)
	}
	go dowaiter(barista,waiter)

	for x := range waiter {
		container = append(container, x)
	}

	return
}

func docashier(vol int,out chan<- string) {
	for x:=1;x<=vol;x++{
		time.Sleep(5*time.Microsecond)
		out<- fmt.Sprintf("order: %d", x)
	}
	close(out)
}

func dobarista(in <-chan string, out chan<- string) {
	for {
		x := <-in
		time.Sleep(100*time.Millisecond)
		out<- fmt.Sprintf("%s %s", x, "espresso")
	}
}

func dowaiter(in <-chan string,out chan<- string) {
	for x := range in {
		time.Sleep(5*time.Millisecond)
		out <- fmt.Sprintf("%s %s", x, "ready :)")
	}
	close(out)
}
