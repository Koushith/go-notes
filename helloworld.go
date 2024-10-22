package main

import "fmt"

func main() {

	var a = "initial" // inferred type
	fmt.Println(a)

	var b, c int = 1, 2 // multiple variables
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // zero value
	fmt.Println(e)

	f := "apple" // shorthand for var f string = "apple"
	fmt.Println(f)
}

// In Go, when you declare a variable without assigning a value, it is automatically initialized with a "zero value" based on its type. This behavior is part of Go's design to ensure that variables always have a well-defined initial state.
// 1. Zero values in Go:
// For numeric types (int, float64, etc.): 0
// For boolean type: false
// For strings: "" (empty string)
// For pointers, functions, interfaces, slices, channels, and maps: nil
// 2. Why Go does this:
// It prevents undefined behavior that can occur with uninitialized variables.
// It provides a consistent and predictable starting point for variables.
// It helps catch potential bugs early by avoiding the use of uninitialized memory.
