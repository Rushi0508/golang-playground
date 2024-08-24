package main

import (
	"fmt"
)

func main() {
	// Using var keyword
	var name string
	name = "Rushi"
	fmt.Println("Name: ", name)

	// Declaring and Initialization together
	var age int = 22
	fmt.Println("Age: ", age)

	// Type Inference
	var city = "Ahmedabad"
	fmt.Println("City: ",city)

	// Shorthand
	state:= "Gujarat"
	fmt.Println("State: ",state)

	// Constants
	const pi = 3.14
	fmt.Println(pi)
	
	// Public Var
	var PublicVar string = "I am accessible outside the package"
	fmt.Println(PublicVar)

	// Private Var: All lower-case starting variables are private
}