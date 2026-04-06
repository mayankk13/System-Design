package main

import (
	"fmt"
	"sync"
)

type Logger struct{}

var (
	instance *Logger
	once     sync.Once
)

func (l *Logger) Log(msg string) {
	fmt.Println(msg)
}

func GetLogger() *Logger {
	// sync.Once is a Go utility
	// It guarantees that a piece of code runs only ONE time
	// Even if multiple goroutines call it at the same time
	once.Do(func() {
		// Creates a new Logger
		// Stores it in instance
		instance = &Logger{}
	})
	return instance
}

func main() {
	l1 := GetLogger()
	l2 := GetLogger()
	/*
			First call
			•	once.Do → runs the function
			•	instance is created

		Second call
			•	once.Do → does NOTHING
			•	Returns existing instance
	*/

	l1.Log("Hello")
	l2.Log("World")

	// Because l1 and l2 point to the exact SAME object in memory.
	fmt.Println("is l1 == l2 ", l1 == l2)
}
