package main

import "fmt"

func main() {
	// Shorthand
	name, email := "Hello world", "test@gmail.com" // type of string
	size := 1.3           // type of float64

	// using var
	var num int32 = 1
	var isbool = true

	// we cannot change the value if using const eg.
	// const isBool = false

	isbool = false
	fmt.Println(name, num, isbool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", num)
	fmt.Printf("%T\n", isbool)
}
