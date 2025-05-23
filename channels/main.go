package main

import "fmt"

func worker(result chan int) {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	result <- sum
}

func main() {
	result := make(chan int)

	go worker(result)

	final := <-result
	fmt.Println("Result:", final)
}
