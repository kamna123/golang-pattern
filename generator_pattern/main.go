package main

import (
	"fmt"
)

/*Generators yields a sequence of values one at a time.*/
func fibonacciGenerator(n int) chan int {
	ch := make(chan int)
	go func() {
		a := 1
		ch <- a
		b := 1
		ch <- b
		c := 2
		for i := 2; i < n; i++ {
			c = a + b
			a = b
			b = c

			ch <- c
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch := fibonacciGenerator(20)
	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(len(ch))
	for i := range ch {
		fmt.Println(i)
	}
}
