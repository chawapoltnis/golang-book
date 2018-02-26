package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
	coins map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InserCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine) SelectSD() string {
	price:=18
	change:=m.insertedMoney-price
	return "SD" + m.change(change)
}

func (m *VendingMachine) SelectCC() string {
	price:=12
	change:=m.insertedMoney-price
	return "CC" + m.change(change)
}

func (m *VendingMachine) change(c int) string{
	if c==0 {
		return ""
	}
	return ", F, TW, O"
}

func main() {
	var coins = map[string]int{"T":10,"F":5,"TW":2,"O":1}
	vm := VendingMachine{coins:coins}
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
}