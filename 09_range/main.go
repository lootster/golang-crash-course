package main

import "fmt"

func main() {
	ids := []int{5,12,34,56,78,90}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	
	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	jobs := map[string]string{"Bob":"Doctor", "Ken":"Pilot,", "Mike":"Teacher"}

	for key, value :=range jobs {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key :=range jobs {
		fmt.Println("Name: ", key)
	}
}
