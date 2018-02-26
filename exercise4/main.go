package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
	coins map[string]int
	items map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InserCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine) SelectSD() string {
	price:=m.items["SD"]
	m.insertedMoney-=price
	return "SD" + m.change(m.insertedMoney)
}

func (m *VendingMachine) SelectCC() string {
	price:=m.items["CC"]
	m.insertedMoney-=price
	return "CC" + m.change(m.insertedMoney)
}

func (m *VendingMachine) change(c int) string{
	var str string
	values := [...]int{10,5,2,1}
	coins := [...]string{"T","F","TW","O"}
	if c>=values[1] {
		str+=", "+coins[1]
		c-=values[1]
	}
	if c>=values[2] {
		str+=", "+coins[2]
		c-=values[2]
	}
	if c>=values[3] {
		str+=", "+coins[3]
		c-=values[3]
	}
	return str
}

func main() {
	var coins = map[string]int{"T":10,"F":5,"TW":2,"O":1}
	var items = map[string]int{"SD":18,"CC":12}
	vm := VendingMachine{coins:coins,items:items}
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