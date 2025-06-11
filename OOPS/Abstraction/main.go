package main

import (
	// "bank"
	"fmt"

	"github.com/mayankk13/System-Design/OOPS/bank"
)

func main() {
	// the user only see the account interface
	// var myAcc bank.Account

	// create a newAccount
	myAcc := bank.NewAccount("Mayank", "XYZ987654", 10000)

	err := myAcc.Deposit(2000)
	if err != nil {
		fmt.Println("Deposit Error:", err)
	}
	fmt.Printf("Balance: ₹%.2f\n", myAcc.GetBalance())

	err = myAcc.Withdraw(5000)
	if err != nil {
		fmt.Println("Withdraw Error:", err)
	}

	fmt.Printf("Remaining Balance: ₹%.2f\n", myAcc.GetBalance())

	// We cannot access internal fields like balance or account number
	// fmt.Println(myAcc.balance) // ❌ will not compile
}
