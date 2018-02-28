package main

import (
	"fmt"
	"github.com/chawapoltnis/vendingMachine"
)

func main() {
	var coins = map[string]int{"T":10,"F":5,"TW":2,"O":1}
	var items = map[string]int{"SD":18,"CC":12}
	vm := vendingMachine.NewVendingMachine(coins,items)
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 0
	vm.InserCoin("T")
	vm.InserCoin("F")
	vm.InserCoin("TW")
	vm.InserCoin("O")
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 18
	can := vm.SelectSD()
	fmt.Println(can) //SD

	vm.InserCoin("T")
	vm.InserCoin("TW")
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 12
	can = vm.SelectCC()
	fmt.Println(can) //CC

	vm.InserCoin("T")
	vm.InserCoin("T")
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 20
	can = vm.SelectCC()
	fmt.Println(can) //CC, F, TW, 0

	vm.InserCoin("T")
	vm.InserCoin("T")
	vm.InserCoin("F")
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 25
	can = vm.CoinReturn()
	fmt.Println(can) //T, T, F
}