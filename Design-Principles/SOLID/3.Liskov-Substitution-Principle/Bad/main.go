package main

import "fmt"

// Account interface
type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	GetBalance() float64
}

// ✅ SavingsAccount works perfectly
type SavingsAcount struct {
	balance float64
}

func (s *SavingsAcount) Deposit(amount float64) {
	s.balance += amount
}

func (s *SavingsAcount) Withdraw(amount float64) {
	s.balance -= amount
}

func (s *SavingsAcount) GetBalance() float64 {
	return s.balance
}

// ❌ FixedDepositAccount violates LSP by panicking on Withdraw
type FixedDepositAccount struct {
	balance float64
}

func (f *FixedDepositAccount) Deposit(amount float64) {
	f.balance += amount
}

func (f *FixedDepositAccount) Withdraw(amount float64) {
	// ❌ This breaks substitutability — throws panic
	panic("Withdraw not allowed on Fixed Deposit")
}

func (f *FixedDepositAccount) GetBalance() float64 {
	return f.balance
}

// Function that uses the Account interface
func PrintAfterWithdraw(a Account) {
	a.Withdraw(1000)
	fmt.Println("Balance:", a.GetBalance())
}

func main() {
	s := &SavingsAcount{}
	s.Deposit(5000)

	f := &FixedDepositAccount{}
	f.Deposit(10000)

	PrintAfterWithdraw(s)
	PrintAfterWithdraw(f)
}
