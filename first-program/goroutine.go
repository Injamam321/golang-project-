package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 12

func printHello(num int) {
	fmt.Println("Hello Imjamam", num)
}

func main() {
	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a, " ", p)

	time.Sleep(5 * time.Second)

}

/*
2.Phases:

	1.Compilation Phase (compile time)
	2.Execution Phase (run time)


***********Compile Phase ********

***Code Segment***


p = 12

printHello = printHello(num int) {...}
main => main() {...}








calculate = func () {...}
calculate anonymous1 = func() {...}
calc = func () {...}
calc anonymous1 = func() {...}
main=func(){...}






go run main.go => compile it => main => ./main
go build main.go => compile it => main

./main

/*
1.slice from an existing array
2.slice from an slice
3.slice literal
4.make function with len
5.make function with len & capacity
6.empty slice or nil slice
7.slice underlying array rule => 1024 -> 100% increase, 1024 < 25% increase
*/
