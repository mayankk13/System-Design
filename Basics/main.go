// given array, arr := []int{11, -3, 5, 8, 12}
// return the square of each number in accending order
// output - [9, 25, 64, 110, 144]
// using concurrency

package main

import (
	"fmt"
	"sort"
	"sync"
)

func square(n int, res *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	num := n * n
	mu.Lock()
	*res = append(*res, num)
	mu.Unlock()
}

func findSquare(arr []int) []int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	n := len(arr)
	res := []int{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go square(arr[i], &res, &wg, &mu)
	}
	wg.Wait()
	return res
}

func main() {
	arr := []int{11, -3, 5, 8, 12}

	res := findSquare(arr)
	sort.Ints(res)

	for i := range res {
		fmt.Println(res[i])
	}
}
