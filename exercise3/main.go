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
	can := vm.selectSD()
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money", vm.GetInsertedMoney())
	can = vm.selectCC()
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money", vm.GetInsertedMoney())
	//change := vm.CoinReturn()
	//fmt.Println(change)
}

func NewVendingMachine() VendingMachine{
	var Ven VendingMachine
	Ven.Coins = map[string]int{"T":10,"F":5,"TW":2,"O":1}
	Ven.Items = map[string]int{"SD":18, "CC":12}
	Ven.TotalCoins=0

	return Ven
}

func (V *VendingMachine)InsertCoin(S string) {
	V.TotalCoins+=V.Coins[S]
}

func (V VendingMachine)GetInsertedMoney() int{
	return V.TotalCoins
}

func (V *VendingMachine)selectSD() string{

	var str string
	P :=[5]int{18,10,5,2,1}
	T :=[5]string{"SD","T","F","TW","O"}
	for i:=0;i<len(P);i++ {
		if V.TotalCoins-P[i]>=0{
			str+=T[i]+" "
			V.TotalCoins-=P[i]
		}
	}
	return str
}

func (V *VendingMachine)selectCC() string{
	var str string
	P :=[5]int{12,10,5,2,1}
	T :=[5]string{"CC","T","F","TW","O"}
	for i:=0;i<len(P);i++ {
		if V.TotalCoins-P[i]>=0{
			str+=T[i]+" "
			V.TotalCoins-=P[i]
		}
	}
	return str
}

func (V *VendingMachine)CoinReturn() string{
	var str string
	P :=[5]int{10,5,2,1}
	T :=[5]string{"T","F","TW","O"}
	for i:=0;i<len(P);i++ {
		if V.TotalCoins/P[i]>1{
			n:=V.TotalCoins/P[i]
			for j:=0;j<n;j++ {
				str+=T[i]+" "
				V.TotalCoins-=P[i]
			}
		}
	}
	return str
}