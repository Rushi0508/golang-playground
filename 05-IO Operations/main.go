package main

import "fmt"

func main() {
    name := "Rushi"
    age := 30
    height := 5.9
    isEmployed := true
    pi := 3.14159
    number := 42

    // Using printf with various format specifiers

    // %s formats a string
    fmt.Printf("Name: %s\n", name)

    // %d formats an integer
    fmt.Printf("Age: %d years\n", age)

    // %.1f formats a float with 1 decimal
    fmt.Printf("Height: %.1f feet\n", height)

    // %t formats a boolean
    fmt.Printf("Employed: %t\n", isEmployed)

     // %.2f formats a float with 2 decimals
    fmt.Printf("PI: %.2f\n", pi)

    // %T formats the type of variable
    fmt.Printf("Number: %d, Type: %T\n", number, number)

    // \n adds a line break
    fmt.Printf("Line 1\nLine 2\n")

	// Using println for basic output

    // Directly outputting strings
    fmt.Println("Name:", name)

    // Directly outputting integers
    fmt.Println("Age:", age, "years")

    // Directly outputting floating-point numbers
    // Note: println does not support formatting decimals
    fmt.Println("Height:", height, "feet")

    // Directly outputting booleans
    fmt.Println("Employed:", isEmployed)

    // Directly outputting floating-point numbers
    // Note: println does not support formatting decimals
    fmt.Println("PI:", pi)

    // Directly outputting integer and type
    // Note: println does not support type formatting
    fmt.Println("Number:", number, "Type:", "int")

    // Directly outputting multi-line text
    // Note: println does not support \n, so use two calls
    fmt.Println("Line 1")
    fmt.Println("Line 2")
}