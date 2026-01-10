package main

import "fmt"

type Logger struct{}

func (l *Logger) Log(msg string) {
	fmt.Println(msg)
}

func main() {
	// &Logger{} → creates a new Logger object in memory
	// l1 and l2 are pointers
	l1 := &Logger{}
	l2 := &Logger{}

	l1.Log("Hello")
	l2.Log("World")

	fmt.Println("is l1 == l2 ", l1 == l2) // Because l1 and l2 point to two different objects in memory.
	// l1  --->  Logger instance #1 (memory A)
	// l2  --->  Logger instance #2 (memory B)
	// Different memory addresses → not equal
}
