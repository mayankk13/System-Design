package account

import (
	"errors"

	"github.com/mayankk13/System-Design/OOPS/bank"
)

type CurrentAccount struct {
	bank.BankAccount
	OverdraftLimit float64
}

func NewCurrentAccount(holder, number string, initial, limit float64) *CurrentAccount {
	return &CurrentAccount{
		BankAccount: bank.BankAccount{
			AccountHolder: holder,
			AccountNumber: number,
			Balance:       initial,
		},
		OverdraftLimit: limit,
	}
}

func (c *CurrentAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	if amount > c.Balance+c.OverdraftLimit {
		return errors.New("overdraft limit exeeded")
	}

	c.Balance -= amount
	return nil
}
