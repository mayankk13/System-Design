package main

import "fmt"

type Account interface {
	Deposit(amount float64)
	GetBalance() float64
	Withdraw(amount float64)
}

type SavingAcount struct {
	balance float64
}

type FixedDepositAccount struct {
	balance float64
}

func (s *SavingAcount) Deposit(amount float64) {
	s.balance += amount
}

func (s *SavingAcount) GetBalance() float64 {
	return s.balance
}

func (s *SavingAcount) Withdraw(amount float64) {
	s.balance -= amount
}

func (f *FixedDepositAccount) Deposit(amount float64) {
	f.balance += amount
}

func (f *FixedDepositAccount) GetBalance() float64 {
	return f.balance
}

func (f *FixedDepositAccount) Withdraw(amount float64) {
	f.balance -= amount
}

func PrintAfterWithdraw(a Account) {
	a.Withdraw(1000)
	fmt.Println("Balance: ", a.GetBalance())
}

func main() {
	savings := &SavingAcount{}
	savings.Deposit(5000)

	fixed := &FixedDepositAccount{}
	fixed.Deposit(10000)

	PrintAfterWithdraw(savings)
	PrintAfterWithdraw(fixed)
}
