package main

import "fmt"

func main() {
	// Create map using make
	scoreBoard := make(map[string]int)
	scoreBoard["John"] = 100
	scoreBoard["Doe"] = 200

	fmt.Println(scoreBoard)

	// Create map using map literal
	scoreBoard2 := map[string]int{
		"John": 100,
		"Doe":  200,
	}

	fmt.Println(scoreBoard2)

	// Accessing map
	fmt.Println(scoreBoard["John"])
	fmt.Println(scoreBoard2["Doe"])

	// Check if key exists
	value, ok := scoreBoard["John"]
	if ok {
		fmt.Println("John's score is", value)
	} else {
		fmt.Println("John's score not found")
	}

	// Delete key
	delete(scoreBoard, "John")
	fmt.Println(scoreBoard)

	// Iterating map
	for key, value := range scoreBoard2 {
		fmt.Println(key, value)
	}
}