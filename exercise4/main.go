package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InserCoin(coin string) {
	if coin=="T" {
		m.insertedMoney+=10
	}
	if coin=="F" {
		m.insertedMoney+=5
	}
	if coin=="TW" {
		m.insertedMoney+=2
	}
	if coin=="O" {
		m.insertedMoney+=1
	}
}

func main() {
	vm := VendingMachine{}
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 0
	vm.InserCoin("T")
	vm.InserCoin("F")
	vm.InserCoin("TW")
	vm.InserCoin("O")
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 10
}