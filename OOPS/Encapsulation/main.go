/*
Encapsulation is one of the core principles of Object-Oriented Programming (OOP). It means hiding internal object
details and exposing only what is necessary via public methods. This prevents external code from directly accessing
internal fields or modifying the state in unintended ways.

Go doesn’t have classes like Java/C++. But it has structs and methods, and it uses export rules
(uppercase for public, lowercase for private) to implement encapsulation.

Feature -
Private fields	-	Start with lowercase (e.g., balance)
Public methods	-	Start with uppercase (e.g., Deposit())
Getter/Setter	-	Use methods to access/update internal data
Access control	-	Not using access modifiers like private/public, but naming convention

Benefit:
Controlled access	-	Can’t change balance directly, must use Deposit() or Withdraw()
Validation			-	Prevents invalid operations (like negative deposits)
Security			-	Protects sensitive fields like balance
Future-proofing		-	Can change internal logic without affecting external code

*/

package main

import (
	"fmt"

	"github.com/mayankk13/System-Design/OOPS/bank"
)

func main() {

	// creating a new bank account
	account := bank.NewAccount("Mayank", "1234567890", 5000)

	fmt.Println("Account holder: ", account.GetAccountHolder())
	fmt.Println("Account number: ", account.GetAccountNumber())
	fmt.Printf("Initial balance: ₹%.2f\n", account.GetBalance())

	// deposit money
	err := account.Deposit(1500)
	if err != nil {
		fmt.Println("Deposit Error:", err)
	} else {
		fmt.Printf("After Deposit: ₹%.2f\n", account.GetBalance())
	}

	// withdraw money more than balance
	err = account.Withdraw(7000)
	if err != nil {
		fmt.Println("Withdraw Error: ", err)
	} else {
		fmt.Printf("After Withdrawal: ₹%.2f\n", account.GetBalance())
	}

	// Valid withdrawal
	err = account.Withdraw(2000)
	if err != nil {
		fmt.Println("Withdraw Error:", err)
	} else {
		fmt.Printf("After Withdrawal: ₹%.2f\n", account.GetBalance())
	}

	// Trying to access balance directly (will not compile)
	// fmt.Println(account.balance) // ❌ invalid, balance is unexported
}
