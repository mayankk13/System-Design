// Your code should be open for extension (via new types) but closed for modification of existing code.

package main

import "fmt"

// Account interface with GetDetails method
type Account interface {
	Deposit(amount float64)
	GetDetails() string
}

// SavingsAccount struct
type SavingAcount struct {
	balance float64
}

func (s *SavingAcount) Deposit(amount float64) {
	s.balance += amount
}

func (s *SavingAcount) GetDetails() string {
	return fmt.Sprintf("Savings Account Balance: ₹%.2f", s.balance)
}

// CurrentAccount struct
type CurrentAccount struct {
	balance float64
}

func (c *CurrentAccount) Deposit(amount float64) {
	c.balance += amount
}

func (c *CurrentAccount) GetDetails() string {
	return fmt.Sprintf("Current Account Balance: ₹%.2f", c.balance)
}

func PrintAccountDetails(a Account) {
	fmt.Println(a.GetDetails())
}

func main() {
	s := &SavingAcount{}
	c := &CurrentAccount{}

	s.Deposit(5000)
	c.Deposit(1500)

	PrintAccountDetails(s)
	PrintAccountDetails(c)
}
