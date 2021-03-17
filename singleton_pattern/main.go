package main

import (
	"fmt"
	"sync"
)

/*the singleton pattern is a software design pattern that
restricts the instantiation of a class to one object.*/

type singleton struct {
	val int
}

var instance *singleton

/*1) Not thread safe(NTS)*/
func GetSingletonNTS() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

/*2) Mutex lock*/
//1) By using mutex locking, you have created an unnecessary bottleneck in your entire program.
var lock *sync.Mutex

func GetSingletonML() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance

}

/* 3)check-lock-check*/
func GetSingletonCLC() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

var once sync.Once

/*4) The right way to implement a singleton pattern in Go is to use the sync package’s Once.Do() function.
This function makes sure that your specified code is executed only once and never more than once.*/
func GetSingletonOnceDo() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	instance := GetSingletonOnceDo()
	instance.val = 10
	fmt.Println(instance.val)
	instance1 := GetSingletonOnceDo()
	fmt.Println(instance1.val)
}
