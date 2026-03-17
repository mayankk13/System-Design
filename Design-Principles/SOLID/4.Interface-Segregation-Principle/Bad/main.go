/*
It states that Interface should be such that the client should NOT implement unnecessary function that
they do not need.
*/

package main

import "fmt"

type Worker interface {
	Work()
	Eat()
	Sleep()
}

type Human struct{}

type Robot struct{}

func (h Human) Work() {
	fmt.Println("Human Working")
}

func (h Human) Eat() {
	fmt.Println("Human Eating")
}

func (h Human) Sleep() {
	fmt.Println("Human Sleeping")
}

func (r Robot) Work() {
	fmt.Println("Robot Working")
}

func (r Robot) Eat() {
	panic("Robot cannot eat")
}

func (r Robot) Sleep() {
	panic("Robot cannot sleep")
}

func StartWork(w Worker) {
	w.Work()
}

func LuncBreak(w Worker) {
	w.Eat()
}

func EndDay(w Worker) {
	w.Sleep()
}

func main() {
	human := Human{}
	robot := Robot{}

	fmt.Println("---- Human Day ----")
	StartWork(human)
	LuncBreak(human)
	EndDay(human)

	fmt.Println("---- Robot Day ----")
	StartWork(robot)
	LuncBreak(robot)
	EndDay(robot)
}
