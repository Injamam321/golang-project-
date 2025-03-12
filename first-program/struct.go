// class -27 struct & receiver function

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user1 User

	user1 = User{
		Name: "Imjamam",
		Age:  30,
	}

	fmt.Println("Name: ", user1.Name)
	fmt.Println("Age: ", user1.Age)

	user2 := User{
		Name: "Roki",
		Age:  16,
	}

	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age: ", user2.Age)

}

/*
2.Phases:

	1.Compilation Phase (compile time)
	2.Execution Phase (run time)


***********Compile Phase ********

***Code Segment***

User = type User struct{...}
main = func () {....}







*/
