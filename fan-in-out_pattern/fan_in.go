package main

import (
	"fmt"
	"sync"
)

func generator(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func squareWorker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		result := num * num // Square the input
		out <- result
	}
}

func main() {
	numbers := generator(1, 5)
	squaredNumbers := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go squareWorker(numbers, squaredNumbers, &wg)
	}

	go func() {
		wg.Wait()
		close(squaredNumbers)
	}()

	for res := range squaredNumbers {
		fmt.Println(res)
	}
}
