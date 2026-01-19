package main

import (
	"fmt"
	"sync"
)

func printNum(i int, ch []chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch[i]
	fmt.Println(i)

	if i < 10 {
		ch[i+1] <- true
	}
}

func main() {
	var wg sync.WaitGroup
	n := 10

	ch := make([]chan bool, n+1)
	for i := 0; i <= n; i++ {
		ch[i] = make(chan bool)
	}

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go printNum(i, ch, &wg)
	}

	ch[1] <- true

	wg.Wait()
}
