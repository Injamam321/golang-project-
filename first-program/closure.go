// class -26 closure

package main

import "fmt"

const a = 10 //constant

var p = 100 //run time

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("==Bank==")
}

/*
2.Phases:

	1.Compilation Phase (compile time)
	2.Execution Phase (run time)


***********Compile Phase ********

***Code Segment***

a = 10
outer = func (){...}
outerAnonymous1==> func(){...}
call = func(){...}
main = func(){...}
init = func(){....}







*/
