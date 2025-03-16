// class -27 struct & receiver function

package main

import "fmt"

type User struct {
	Name string //member variable or property
	Age  int
}

func printUserdetails(usr User) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)

}

//Receiver Function
func (usr User) printdetails() {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)

}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	var user1 User

	user1 = User{ //instantiate
		Name: "Imjamam",
		Age:  30,
	}

	// printUserdetails(user1)
	user1.printdetails()
	user1.call(20)

	user2 := User{ //instance or object
		Name: "Roki",
		Age:  16,
	}
	printUserdetails(user2)

}

/*
2.Phases:

	1.Compilation Phase (compile time)
	2.Execution Phase (run time)


***********Compile Phase ********

***Code Segment***

User = type User struct{...}
printUserdetails = func  (usr User){..}
printdetails = func() {...} //User type only call that func
call = func (a int) {...} //User type only call that func
main = func () {....}





go run reciver.go => compile it => reciver => ./reciver
go build reciver.go => compile it => reciver


./reciver


*/
