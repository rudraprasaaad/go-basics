/*
limited number of go routines doing unlimited jobs
worker pool will take one by one from a channel until no jobs are left
*/
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d: finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}

}
