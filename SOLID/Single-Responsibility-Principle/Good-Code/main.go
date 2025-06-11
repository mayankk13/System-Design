package main

import (
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
}

// handles buisness logic
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

// separate type for report generation
type ReportGenerator struct{}

func (r *ReportGenerator) GenerateReport(a Account) string {
	return fmt.Sprintf("Account: %s, Balance: ₹%.2f", a.AccountNumber, a.Balance)
}

func main() {
	a := Account{
		AccountNumber: "1234",
		Balance:       5000,
	}

	a.Deposit(1500)

	reportor := ReportGenerator{}
	fmt.Println(reportor.GenerateReport(a))
}
