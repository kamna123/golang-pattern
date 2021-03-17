package main

import (
	"fmt"
	"time"
)

/*Decorators essentially allow you to wrap existing functionality and append or prepend
your own custom functionality on top.*/
func mainFun() {
	fmt.Println("main func")
	time.Sleep(1 * time.Second)
}

func additionalFun(a func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("ending function execution: %s\n", time.Now())
}
func main() {
	additionalFun(mainFun)
}
