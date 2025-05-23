package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	ch <- 25
	ch <- 30
	ch <- 35

	data := <-ch
	fmt.Println("Received 1:", data)
	data = <-ch
	fmt.Println("Received 2:", data)
	data = <-ch
	fmt.Println("Recieved 3:", data)
}

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg, ch)
	wg.Wait()
}
