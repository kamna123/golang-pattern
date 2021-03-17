package main

import "fmt"

/*In Java, you have to explicitly tell the class to implements defined interface in
the concrete class, so Java compiler can understand your intention,
but in Go it’s done implicitly.*/

type Engine interface {
	start()
	stop()
}
type car struct {
}
type train struct {
}

func (c car) start() {
	fmt.Println("My car engine is started")
}
func (c car) stop() {
	fmt.Println("My car engine is stopped")
}

func (t train) start() {
	fmt.Println("My train engine is started")
}
func (t train) stop() {
	fmt.Println("My train engine is stopped")
}
func starting(e Engine) {
	e.start()
}

func stopping(e Engine) {
	e.stop()
}
func main() {
	carEngine := car{}
	trainEngine := train{}
	engines := []Engine{carEngine, trainEngine}
	for _, engine := range engines {
		starting(engine)
		stopping(engine)
	}
}
