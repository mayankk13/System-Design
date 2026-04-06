package main

import "fmt"

// concrete struct
type CardPayment struct{}

func (c CardPayment) Pay(amount int) {
	fmt.Println("Paid using Card: ", amount)
}

type UPIPayment struct{}

func (u UPIPayment) Pay(amount int) {
	fmt.Println("Paid using UPI: ", amount)
}

func main() {
	// client decides which payment to create
	method := "upi"

	var payment interface{}

	if method == "card" {
		payment = CardPayment{}
	} else if method == "upi" {
		payment = UPIPayment{}
	}

	if p, ok := payment.(UPIPayment); ok {
		p.Pay(500)
	}
}
