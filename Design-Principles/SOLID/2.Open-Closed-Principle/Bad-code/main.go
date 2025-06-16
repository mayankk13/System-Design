// This approach uses conditional logic that needs to be modified every time a new account type is introduced.

package main

import "fmt"

type Account struct {
	AccountType string
	Balance     float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

// this func breaks OCP
func PrintAccountDetails(a Account) {
	if a.AccountType == "Savings" {
		fmt.Printf("Saving account balance: ₹%0.2f\n", a.Balance)
	} else if a.AccountType == "Current" {
		fmt.Printf("Current account balance: ₹%0.2f\n", a.Balance)
	} else {
		fmt.Println("Unknown account type")
	}
}

func main() {
	acc1 := Account{
		AccountType: "Savings",
		Balance:     5000,
	}

	acc2 := Account{
		AccountType: "Current",
		Balance:     1500,
	}

	PrintAccountDetails(acc1)
	PrintAccountDetails(acc2)
}

/*
Problem:
	•	If you add a new account type (like LoanAccount), you have to modify PrintAccountDetails.
	•	This violates the Open/Closed Principle — code is not closed to modification.
*/
