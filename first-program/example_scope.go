// Package-level variables
var (
	a = 10
	b = 20
)

func add() {
	fmt.Println("Sum:", a+b) // a এবং b এখানে package-level থেকে অ্যাক্সেস করা হচ্ছে
}

func main() {
	add() // ফাংশন add() call করলে package-level ভেরিয়েবল a এবং b ব্যবহার করবে
}



// Function-Level Scope (Local Scope)

package main

import "fmt"

func calculateSum() {
	x := 15 // Local variable
	y := 25 // Local variable
	fmt.Println("Sum:", x+y)
}

func main() {
	calculateSum()
	// fmt.Println(x) // এই লাইনটি error দেবে কারণ x ফাংশন-লেভেল স্কোপের বাইরে
}


// Block Scope (Nested Scope)

package main

import "fmt"

func main() {
	if true {
		x := 50 // x এর scope শুধুমাত্র এই ব্লকের মধ্যে
		fmt.Println("Inside if block:", x)
	}
	// fmt.Println(x) // Error, কারণ x if ব্লকের বাইরে অ্যাক্সেসযোগ্য নয়
}


// Combining Multiple Scopes:

package main

import "fmt"

// Package-level variable
var globalVar = "I am a global variable"

func main() {
	localVar := "I am a local variable" // Function-level variable

	fmt.Println(globalVar) // Accessing package-level variable
	fmt.Println(localVar)  // Accessing function-level variable

	for i := 0; i < 3; i++ {
		blockVar := i * 2 // Block-level variable
		fmt.Println("Inside loop:", blockVar)
	}
	// fmt.Println(blockVar) // Error, blockVar এর scope শুধুমাত্র loop ব্লকের মধ্যে
}
