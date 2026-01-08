package main

import (
	"fmt"
	"sync"
)

func printHello(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello ")
	ch <- true
}

func printWorld(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch
	fmt.Printf("World!")
}

func main() {
	ch := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)

	go printHello(ch, &wg)
	go printWorld(ch, &wg)

	wg.Wait()
}
