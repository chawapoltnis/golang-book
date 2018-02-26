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
	m.insertedMoney=0
	return "SD"
}

func (m *VendingMachine) SelectCC() string {
	m.insertedMoney=0
	return "CC"
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
	fmt.Println(can)

	vm.InserCoin("T")
	vm.InserCoin("TW")
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 12
	can = vm.SelectCC()
	fmt.Println(can)
}