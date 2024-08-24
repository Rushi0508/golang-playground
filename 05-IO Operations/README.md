## IO Operations

### How to output?

In Go, the `Printf` and `Println` functions are used for formatted and simple output, respectively.

#### `Printf` Example

The `Printf` function in Go is used for formatted output, where you can specify how the data should be displayed.

```go
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
    fmt.Printf("Line 1\nLine 2")
}
```

Output

```go
Name: Rushi
Age: 30 years
Height: 5.9 feet
Employed: true
PI: 3.14
Number: 42, Type: int
Line 1
Line 2
```

#### `Println` Example

The `println` function in Go is used for basic output. It automatically separates arguments with spaces and adds a newline after each call. Unlike `printf`, `println` does not support format specifiers.

Here we have used the same example for better understanding

```go
package main

import "fmt"

func main() {
    name := "Rushi"
    age := 30
    height := 5.9
    isEmployed := true
    pi := 3.14159
    number := 42

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
```

### How to input?

There are various methods for taking input like `Scan`, `Scanf`, `Scanln` of `fmt` package and `NewReader` of `bufio` package

1. Using `Scan`

- The `fmt.Scan` function reads space-separated input values and stores them in provided variables.

```go
var name string
var age int

fmt.Print("Enter your name and age: ")
fmt.Scan(&name, &age)
```

2. Using `Scanf`

- The `fmt.Scanf` function reads input according to a specified format. It allows for more control over input parsing.

```go
var name string
var age int
var height float64

fmt.Print("Enter your name, age, and height (space-separated): ")
fmt.Scanf("%s %d %f", &name, &age, &height)
```

3. Using `Scanln`

- The fmt.Scanln function reads input until a newline is encountered. It is useful for reading entire lines of input.

```go
var name string
var age int

fmt.Print("Enter your name: ")
fmt.Scanln(&name)

fmt.Print("Enter your age: ")
fmt.Scanln(&age)
```

4. Using `NewReader`

- The `bufio` package provides a `Reader` that allows you to read input line by line, which is useful for handling user input more flexibly.

```go
fmt.Println("Enter your name")

reader := bufio.NewReader(os.Stdin)

// This line reads a line of input from the user until it encounters a newline character (\n).
// The second value (which is ignored here with _) is an error value that is returned by ReadString.
name, _ := reader.ReadString('\n')
```
