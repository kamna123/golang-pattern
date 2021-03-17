package main

import (
	"fmt"
	"sync"
)

/*https://dev.to/nadirbasalamah/golang-tutorial-10-concurrency-with-channel-54co*/
/**/

func main() {
	ints := generateNum()
	c1 := average(ints)
	c2 := average(ints)
	for val := range merge(c1, c2) {
		fmt.Println(val)
	}
}
func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ch))
	output := func(c <-chan int) {
		for val := range c {
			out <- val
		}
		wg.Done()
	}
	for _, val := range ch {
		go output(val)

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func generateNum() <-chan []int {
	out := make(chan []int)
	go func() {
		data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		for _, val := range data {
			fmt.Println("inserting ", val)
			out <- val
		}
		close(out)
	}()
	return out
}

func average(i <-chan []int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range i {
			num := avg(val)
			fmt.Println("calculating ", num)
			out <- num
		}
		close(out)
	}()
	return out
}

func avg(num []int) int {
	sum := 0
	for _, val := range num {
		sum = sum + val
	}
	return sum / len(num)
}
