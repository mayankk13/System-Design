// “Keep your code as simple as possible. Avoid unnecessary complexity.”

package main

import "fmt"

type NumberChecker struct{}

func (nc *NumberChecker) IsEven(number int) (bool, error) {
	if number < 0 {
		return false, fmt.Errorf("negative numbers are not supported") // 🚨 unnecessary logic
	}

	if number%2 == 0 {
		return true, nil
	}
	return false, nil
}

func main() {
	nc := &NumberChecker{}
	even, err := nc.IsEven(4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Is even:", even)
}

/*
What’s wrong?
	•	Too much code for a very simple task.
	•	Why do we need a struct? Why are we returning an error?
Just to check if a number is even?
*/
