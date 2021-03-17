package main

import "fmt"

/* 1. Use factory functions to ensure that new instances structs are
constructed with the required arguments
2. Factory Method defines a method, which should be used for creating objects
instead of direct constructor call (new operator). Subclasses can override
this method to change the class of objects that will be created.
3. Factory method creational design pattern allows creating objects without
having to specify the exact type of the object that will be created.
4. Hide detail of creational object, and delegate the instantiation 
(even dependency injection) to the factory
5. Decouple the abstraction and implementation, 
so factory can do the creational operation
6. Example : Let’s say, our previous Car and Train are produced in same factory, 
we call it as Engine Factory. The boss of the factory doesn’t need to know the detail 
implementation of how the Car Engine or Train Engine is produced. The boss only send a 
message to the factory manager, “Bring me the car engine”, then the factory manager will 
decide to gow where the assembly line of car engine is located. After factory manager 
get the wanted engine, he will come back to the boss with that engine.*/

type Engine interface{
	Start()
	Stop()
}
type car struct{

}
func(c car) Start(){
	fmt.Println("car start")
}
func(c car) Stop(){
	fmt.Println("car stop")
}

type train struct{

}
func(c train) Start(){
	fmt.Println("train start")
}
func(c train) Stop(){
	fmt.Println("train stop")
}

func starting( e Engine){
	e.Start()
}

func stopping( e Engine){
	e.Stop()
}

func GetEngine(engineType string) Engine {
	switch engineType {
	case "car":
		return car{}
	case "train":
		return train{}
	default:
		fmt.Println("type undefined")
		return nil
	}
}

func main(){
	engine := GetEngine("car")

	starting(engine)
	stopping(engine)
	engine1 := GetEngine("train")
	starting(engine1)
	stopping(engine1)
}