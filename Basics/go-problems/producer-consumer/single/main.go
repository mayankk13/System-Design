package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced - ", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Consumed - ", val)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}
