package main

import "fmt"

/*Builder pattern separates the construction of a complex object from its representation
so that the same
construction process can create different representations.*/

type User struct {
	Name   string
	Role   string
	Salary int
}
type UserBuilder struct {
	User
}

func (ub *UserBuilder) setName(name string) *UserBuilder {
	ub.User.Name = name
	return ub
}

func (ub *UserBuilder) setRole(role string) *UserBuilder {
	ub.User.Role = role
	return ub
}

func (ub *UserBuilder) setSalary(sal int) *UserBuilder {
	ub.User.Salary = sal
	return ub
}
func (ub *UserBuilder) Build() User {
	return ub.User
}
func main() {

	ub := &UserBuilder{}
	user := ub.
		setName("Michael Scott").
		setRole("manager").
		setSalary(1000).Build()

	fmt.Println(user)
}
