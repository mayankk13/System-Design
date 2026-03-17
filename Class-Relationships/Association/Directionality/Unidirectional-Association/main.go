package main

import "fmt"

// PaymentGatway has a method to process Payment
// It does not know anything about order
type PaymentGatway struct{}

func (pg *PaymentGatway) processPayment(amount float64) {
	fmt.Println("Processing Payment: ", amount)
}

// order depend on PaymentGateway
// Order --> PaymentGateway (unidirectional)

type Order struct {
	gateway *PaymentGatway // * because You want to share the same object, not a copy
}

func NewOrder(gateway *PaymentGatway) *Order {
	return &Order{gateway: gateway}
}

func (o *Order) Checkout() {
	o.gateway.processPayment(100.0)
}

func main() {
	gateway := &PaymentGatway{}
	order := NewOrder(gateway)

	order.Checkout()
}
