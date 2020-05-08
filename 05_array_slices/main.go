package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Declare and assign
	vegeArr := [2]string{"Carrot", "Spinach"}

	fmt.Println(vegeArr)
	fmt.Println(vegeArr[1])

	vegeSlice := []string{"Carrot", "Spinach", "Cucumber"}
	fmt.Println(vegeSlice)

	// Count length of array
	fmt.Println(len(vegeSlice))

	// Slicing array
	fmt.Println(vegeSlice[1:2])
}
