/*
“Every piece of logic should be written in one place only.”
If you copy-paste the same logic in multiple places, you violate DRY.

DRY makes code easier to maintain, reuse, and test.
*/

package main

import "fmt"

type Order struct {
	Amount float64
}

func applyDicount(order Order, channel string) float64 {
	discount := order.Amount * 0.1
	finalAmount := order.Amount - discount

	if channel == "online" {
		fmt.Println("🛒 Online Order Final Amount:", finalAmount)
	} else if channel == "in-store" {
		fmt.Println("🏬 In-Store Order Final Amount:", finalAmount)
	}

	return finalAmount
}

func main() {
	order1 := Order{Amount: 1000}
	order2 := Order{Amount: 1500}

	applyDicount(order1, "online")
	applyDicount(order2, "in-store")
}
