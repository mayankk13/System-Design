package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started the job %d\n", id, job)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished the job %d\n", id, job)
	}
}

func main() {
	numberOfWorkers := 3
	numberOfJobs := 10

	jobs := make(chan int)
	var wg sync.WaitGroup

	// start workers
	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	// allocate jobs
	for i := 1; i <= numberOfJobs; i++ {
		jobs <- i
	}

	close(jobs)

	fmt.Println("All job finished")
}
