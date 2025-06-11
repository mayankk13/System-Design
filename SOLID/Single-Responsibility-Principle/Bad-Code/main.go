package main

import "fmt"

type Account struct {
	AccountNumber string
	Balance       float64
}

// Handles business logic
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

// also handles report generation (UI logic)
func (a *Account) GenerateReport() string {
	return fmt.Sprintf("Account: %s, Balance: ₹%.2f", a.AccountNumber, a.Balance)
}

func main() {
	a := Account{
		AccountNumber: "1234",
		Balance:       5000,
	}

	a.Deposit(1500)
	fmt.Println(a.GenerateReport())
}
