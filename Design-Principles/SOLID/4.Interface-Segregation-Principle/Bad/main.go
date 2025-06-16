package main

import "fmt"

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	RequestLoan(amount float64)
}

type FixedDepositAccount struct {
	Balance float64
}

func (f *FixedDepositAccount) Deposit(amount float64) {
	f.Balance += amount
	fmt.Println("✅ FD: Deposit successful")
}

func (f *FixedDepositAccount) Withdraw(amount float64) {
	panic("❌ FD: Withdraw not allowed") // Bad: panic if someone calls it
}

func (f *FixedDepositAccount) RequestLoan(amount float64) {
	panic("❌ FD: Loan not allowed") // Bad: panic again
}

func main() {
	account := &FixedDepositAccount{}

	account.Deposit(5000)

	// These should NOT be accessible!
	account.Withdraw(2000)
	account.RequestLoan(40000)
}
