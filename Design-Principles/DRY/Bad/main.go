package main

import "fmt"

type Order struct {
	Amount float64
}

func ApplyDicountForOnline(order Order) float64 {
	discount := order.Amount * 0.1
	finalAmount := order.Amount - discount
	fmt.Println("🛒 Online Order Final Amount:", finalAmount)

	return finalAmount
}

func ApplyDicountForInStore(order Order) float64 {
	discount := order.Amount * 0.1
	finalAmount := order.Amount - discount
	fmt.Println("🛒 Online Order Final Amount:", finalAmount)

	return finalAmount
}

func main() {
	order1 := Order{Amount: 1500}
	order2 := Order{Amount: 1000}

	ApplyDicountForOnline(order1)
	ApplyDicountForInStore(order2)
}

/*
 What’s wrong?
	•	The same discount logic is written twice: once for online, once for in-store.
	•	If discount changes (say from 10% to 15%), you’ll have to change it everywhere — risk of bugs.
*/
