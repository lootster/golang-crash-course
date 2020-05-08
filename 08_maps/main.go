package main

import "fmt"

func main() {
	// Define maps
	emails := make(map[string]string)

	// Assign key value
	emails["Bob"] = "bob@gmail.com"
	emails["Ken"] = "ken@hotmail.com"
	emails["Mike"] = "mike@yahoo.com"

	fmt.Println(emails)

	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])
	
	// Delete Bob
	delete(emails, "Bob")
	fmt.Println(emails)
	
	// Declare map and add key value
	jobs := map[string]string{"Bob":"Doctor", "Ken":"Pilot,", "Mike":"Teacher"}
	fmt.Println(jobs)

}
