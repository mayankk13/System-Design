package main

import "fmt"

// step 1 - interface
type Payment interface {
	Pay(amount int)
}

// step 2 - implementation
type CardPayment struct{}

func (c CardPayment) Pay(amount int) {
	fmt.Println("Paid using card: ", amount)
}

type UPIPayment struct{}

func (u UPIPayment) Pay(amount int) {
	fmt.Println("Paid using UPI: ", amount)
}

// step 3 - factory method
func PaymentFactory(method string) Payment {
	switch method {
	case "card":
		return &CardPayment{}
	case "upi":
		return &UPIPayment{}
	default:
		return nil
	}
}

func main() {
	payment := PaymentFactory("upi")

	if payment == nil {
		fmt.Println("invalid payment method!")
		return
	}

	payment.Pay(1000)
}
