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

func (u *UserBuilder) setName(name string) *UserBuilder {
	u.User.Name = name
	return u
}

func (u *UserBuilder) setRole(role string) *UserBuilder {
	u.User.Role = role
	return u
}

func (u *UserBuilder) setSalary(sal int) *UserBuilder {
	u.User.Salary = sal
	return u
}
func (ub *UserBuilder) Build() User {
	return ub.User
}
func main() {

	/*The value of a constant should be known at compile time. Hence it cannot
	be assigned to a value returned by a function call since the function call
	takes place at run time.*/

	//	var a = math.Sqrt(4)   //allowed
	//  const b = math.Sqrt(4) //not allowed
	ub := &UserBuilder{}
	user := ub.
		setName("Michael Scott").
		setRole("manager").
		setSalary(1000).Build()

	fmt.Println(user)
}
