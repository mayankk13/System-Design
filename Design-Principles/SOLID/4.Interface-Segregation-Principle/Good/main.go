package main

import "fmt"

// small, focus interfaces
type Depositable interface {
	Deposit(amount float64)
}

type Withdrawable interface {
	Withdraw(amount float64)
}

type LoanRequestable interface {
	RequestLoan(amount float64)
}

// --- Fixed deposit supports only deposit
type FixedDepositAccount struct {
	Balance float64
}

func (f *FixedDepositAccount) Deposit(amount float64) {
	f.Balance += amount
	fmt.Println("✅ FD: Deposit successful")
}

// --- Savings supports deposit, withdraw, and loan
type SavingsAccount struct {
	Balance float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.Balance += amount
	fmt.Println("✅ Savings: Deposit successful")
}

func (s *SavingsAccount) Withdraw(amount float64) {
	s.Balance -= amount
	fmt.Println("✅ Savings: Withdraw successful")
}

func (s *SavingsAccount) RequestLoan(amount float64) {
	fmt.Println("✅ Savings: Loan request submitted")
}

// --- Functions using only what they need

func ProcessDeposit(d Depositable) {
	d.Deposit(10000)
}

func ProcessWithdrawal(w Withdrawable) {
	w.Withdraw(2000)
}

func ProcessLoanRequest(l LoanRequestable) {
	l.RequestLoan(50000)
}

func main() {
	fixed := &FixedDepositAccount{}
	savings := &SavingsAccount{}

	ProcessDeposit(fixed)   // ✅
	ProcessDeposit(savings) // ✅

	ProcessWithdrawal(savings)  // ✅
	ProcessLoanRequest(savings) // ✅

	// ❌ Compile error if uncommented — and that's good
	// ProcessWithdrawal(fixed)
	// ProcessLoanRequest(fixed)
}
