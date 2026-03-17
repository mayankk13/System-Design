package main

import "fmt"

type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

type Sleeper interface {
	Sleep()
}

type Human struct{}

type Robot struct{}

// HUman implements all
func (h Human) Work() {
	fmt.Println("Human is Working")
}

func (h Human) Eat() {
	fmt.Println("Human is Eating")
}

func (h Human) Sleep() {
	fmt.Println("Human is Sleeping")
}

// Robot implements only what it needs
func (r Robot) Work() {
	fmt.Println("Robot is working")
}

func StartWork(w Worker) {
	w.Work()
}

func LuncBreak(e Eater) {
	e.Eat()
}

func EndDay(s Sleeper) {
	s.Sleep()
}

func main() {
	human := Human{}
	robot := Robot{}

	fmt.Println("--- Human Day ---")
	StartWork(human)
	LuncBreak(human)
	EndDay(human)

	fmt.Println("--- Robot Day ---")
	StartWork(robot)
}
