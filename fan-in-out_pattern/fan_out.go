package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func worker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		result := num * 2 // Process the input
		out <- result
	}
}

func main() {
	input := make(chan int)
	output := make(chan int)

	go producer(input)

	var wg sync.WaitGroup
	workerCount := 3
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go worker(input, output, &wg)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for res := range output {
		fmt.Println(res)
	}
}
