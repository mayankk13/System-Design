// creational pattern
package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func NewUser(name, email string, age int) User {
	if name == "" {
		panic("name cannot be empty - ")
	}

	return User{
		Name:  name,
		Email: email,
		Age:   29,
	}
}

func main() {
	// insted of calling User struct directly -
	// u1 := User{"Mayank", "mayank@gmail.com", 25}

	u1 := NewUser("mayank", "mayank@gmail.com", 29)

	fmt.Println(u1)
}
