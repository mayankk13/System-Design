package main

import (
	"fmt"
	"sync"
)

func printEven(odd, even chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= 10; i = i + 2 {
		<-even
		fmt.Println(i)
		if i < 10 {
			odd <- true
		}
	}
}

func printOdd(odd, even chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 10; i = i + 2 {
		<-odd
		fmt.Println(i)
		even <- true
	}
}

func main() {
	var wg sync.WaitGroup
	even := make(chan bool)
	odd := make(chan bool)

	wg.Add(2)

	go printOdd(odd, even, &wg)
	go printEven(odd, even, &wg)

	odd <- true

	wg.Wait()
}
