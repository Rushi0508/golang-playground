## Conditions in Go (if-else and switch)

- Conditions are used to execute a block of code based on a condition.

### If-Else

- The `if` statement is used to execute a block of code if the condition is true.

- The `else` statement is used to execute a block of code if the condition is false.

- The `else` statement is optional. But it must be used after the `if` or `else if` statement.

- The `else if` statement is used to execute a block of code if the previous condition is false.

```go
package main

import "fmt"

func main() {
    num := 10

    if num > 10 {
        fmt.Println("Number is greater than 10")
    } else if num < 10 {
        fmt.Println("Number is less than 10")
    } else {
        fmt.Println("Number is equal to 10")
    }
}
```

In above example `num` is equal to 10, so the output will be:

```go
Number is equal to 10
```

### Switch

- The `switch` statement is used to execute a block of code based on the value of an expression.

- The `switch` statement is used to replace multiple `if` statements.

- `case` is used to define the value to compare. If the value of the expression is equal to the value of the `case`, then the block of code inside the `case` will be executed.

- `default` is used to execute a block of code if none of the `case` values match the value of the expression.

```go
func main() {
    num := 10

    switch num {
    case 10:
        fmt.Println("Number is 10")
    case 20:
        fmt.Println("Number is 20")
    default:
        fmt.Println("Number is not 10 or 20")
    }
}
```

In above example `num` is equal to 10, so the output will be:

```go
Number is 10
```
