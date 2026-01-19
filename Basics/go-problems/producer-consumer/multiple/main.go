package main

import (
	"fmt"
	"sync"
)

func producer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("producer %d, produced %d\n", id, i)
		ch <- i
	}
}

func consumer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		fmt.Printf("consumer %d, consumed %d\n", id, val)
	}
}

func main() {
	var prodWg sync.WaitGroup
	var consWg sync.WaitGroup

	ch := make(chan int)

	numProducers := 3
	numConsumers := 2

	// start producer
	for i := 0; i < numProducers; i++ {
		prodWg.Add(1)
		go producer(i, ch, &prodWg)
	}

	// start consumer
	for i := 0; i < numConsumers; i++ {
		consWg.Add(1)
		go consumer(i, ch, &consWg)
	}

	go func() {
		prodWg.Wait()
		close(ch)
	}()

	consWg.Wait()

}
