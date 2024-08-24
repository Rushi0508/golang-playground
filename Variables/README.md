## Variables in Go

### Basic Variable Declaration

There are several ways to declare variables in Go:

1. **Using the `var` keyword:**

    ```go
    var name string
    var age int
    ```

    This declares a variable without initializing it. The zero value of the type will be assigned (`""` for strings, `0` for integers, etc.).

2. **Declaring and initializing together:**

    ```go
    var name string = "John"
    var age int = 30
    ```

3. **Type inference:**

    Go can automatically infer the type based on the value assigned.

    ```go
    var name = "John"
    var age = 30
    ```

4. **Short variable declaration (inside functions only):**

    This is the shorthand syntax for declaring and initializing variables.

    ```go
    name := "John"
    age := 30
    ```

    Note that this syntax cannot be used outside functions.

### Zero Values

When a variable is declared but not initialized, Go assigns a zero value to the variable:

- `0` for numeric types.
- `""` for strings.
- `false` for booleans.
- `nil` for pointers, slices, maps, interfaces, channels, and function types.

Example:

```go
var (
    i int       // 0
    f float64   // 0.0
    b bool      // false
    s string    // ""
)
```

### Multiple Variable Declaration

You can declare and initialize multiple variables at once.

```go
var name, age = "John", 30
```

This assigns the value "John" to name and 30 to age.

### Constants

Constants are declared using the const keyword. They cannot be changed once assigned.

```go
const pi = 3.14
```

### Public Variables

Public variables (or exported variables) are variables whose names start with an uppercase letter. These variables are accessible from other packages.

Example:

```go
package mypackage

// Public variable
var PublicVar string = "I am accessible outside the package"
```

If you import mypackage in another Go file, you can access PublicVar:

```go
package main

import (
    "fmt"
    "mypackage"
)

func main() {
    fmt.Println(mypackage.PublicVar)
}
```

### Private Variables

Private variables are declared by starting the variable name with a lowercase letter:

```go
package mypackage

// Private variable
var privateVar string = "I am only accessible within this package"
```

Attempting to access privateVar from another package will result in a compilation error:

```go
package main

import (
    "fmt"
    "mypackage"
)

func main() {
    // This will cause an error because privateVar is not exported
    fmt.Println(mypackage.privateVar) // ERROR: privateVar is not exported by package mypackage
}
```
