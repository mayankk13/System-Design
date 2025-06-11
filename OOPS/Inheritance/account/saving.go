package account

import (
	"github.com/mayankk13/System-Design/OOPS/bank"
)

type SavingAcount struct {
	bank.BankAccount
	InterestRate float64
}

func NewSAvingAccount(holder, number string, initial, rate float64) *SavingAcount {
	return &SavingAcount{
		BankAccount: bank.BankAccount{
			AccountHolder: holder,
			AccountNumber: number,
			Balance:       initial,
		},
		InterestRate: rate,
	}
}

func (s *SavingAcount) AddInterest() {
	interest := s.Balance * s.InterestRate / 100
	s.Balance += interest
}
