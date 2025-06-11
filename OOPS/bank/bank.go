package bank

import (
	"errors"
)

// step1: define the abstract interface
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
	GetAccountHolder() string
	GetAccountNumber() string
}

// Step2: Private concrete implementation of Account
type BankAccount struct {
	AccountHolder string
	AccountNumber string
	Balance       float64
}

// step3: Constructor to return Account (not bankAccount)
func NewAccount(holder, number string, initial float64) Account {
	if initial < 0 {
		initial = 0
	}

	return &BankAccount{
		AccountHolder: holder,
		AccountNumber: number,
		Balance:       initial,
	}
}

// step4: implement the interface methods
func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit must be positive")
	}

	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw must be positive")
	}

	if amount > b.Balance {
		return errors.New("insufficient fund")
	}

	b.Balance -= amount
	return nil
}

func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b *BankAccount) GetAccountHolder() string {
	return b.AccountHolder
}

func (b *BankAccount) GetAccountNumber() string {
	return b.AccountNumber
}
