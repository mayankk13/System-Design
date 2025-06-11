/*
In classical OOP, inheritance means a class (child) can inherit properties and methods from another class (parent).
Go doesn’t have extends or super, but it provides composition (embedding one struct inside another) to achieve the same result.

We’ll simulate two types of accounts:
	1.	BaseAccount – common properties and methods
	2.	SavingsAccount – extends BaseAccount and adds interest
	3.	CurrentAccount – adds overdraft capability

Feature				Traditional OOP					Go Equivalent

Class inheritance	class B extends A				Embed struct A in struct B
Method override		Override method in subclass		Define method on embedded struct
Polymorphism		Base class references			Use interfaces
*/

package main

import (
	"fmt"

	"github.com/mayankk13/System-Design/OOPS/Inheritance/account"
)

func main() {
	// saving account
	savings := account.NewSAvingAccount("Mayank", "SAV123", 10000, 4.0)
	savings.Deposit(2000)
	savings.AddInterest()
	fmt.Printf("Savings Balance (with interest): ₹%.2f\n", savings.GetBalance())

	// current account
	current := account.NewCurrentAccount("Mayank", "CUR123", 5000, 2000)
	err := current.Withdraw(6500)
	if err != nil {
		fmt.Println("Withdraw failed:", err)
	} else {
		fmt.Printf("Current Balance after overdraft: ₹%.2f\n", current.GetBalance())
	}
}
