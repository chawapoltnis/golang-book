package main

import (
	"fmt"
)

type VendingMachine struct {
	Coins map[string]int
	Items map[string]int
	TotalCoins int
}

func main() {
	//vm := NewVendingMachine(coins, items)
	vm := NewVendingMachine()

	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money", vm.GetInsertedMoney())
	//can := vm.selectSD()
	//fmt.Println(can)
}

func NewVendingMachine() VendingMachine{
	var Ven VendingMachine
	Ven.Coins = map[string]int{"T":10,"F":5,"TW":2,"0":1}
	Ven.Items = map[string]int{"SD":18}
	Ven.TotalCoins=0

	return Ven
}

func (V *VendingMachine)InsertCoin(S string) {
	V.TotalCoins+=V.Coins[S]
}

func (V VendingMachine)GetInsertedMoney() int{
	return V.TotalCoins
}