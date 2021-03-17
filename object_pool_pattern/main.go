package main

import "fmt"

type Pool chan *Object
type Object struct{}

/*The object pool creational design pattern is used to prepare
and keep multiple instances according to the demand expectation.*/
func NewPool(total int) *Pool {
	p := make(Pool, total) // buffered channel of objects

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

func main() {
	TotalChannels := 10
	p := NewPool(TotalChannels)

	for {
		select {
		case <-*p:
			//obj.Do( /*...*/ )
			fmt.Println(" obj consumed")
			*p <- obj
		default:
			// No more objects left — retry later or fail
			return
		}
	}
}
