package main

import (
	"fmt"
)

type VendingMachine struct {

}

func (m VendingMachine) InsertedMoney() int {
	return 10
}

func (m *VendingMachine) InserCoin(coin string) {

}

func main() {
	vm := VendingMachine{}
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 0
	vm.InserCoin("T")
	fmt.Println("Inserted Money:",vm.InsertedMoney()) //Inserted Money: 10
}