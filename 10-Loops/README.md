## Loops in Go

In Go, loops are used to execute a block of code repeatedly.

The `for` loop is the only loop available in Go. The `for` loop is similar to the `for` loop in C and Java.

The `for` loop has three components: the initialization statement, the condition statement, and the post statement. The initialization statement is executed only once at the beginning of the loop. The condition statement is evaluated before each iteration of the loop. If the condition is true, the block of code inside the loop is executed. The post statement is executed after each iteration of the loop.

Here is the syntax of the `for` loop in Go:

```go
for initialization; condition; post {
    // code block to be executed
}
```

The `for` loop can also be used as a `while` loop by omitting the initialization and post statements. Here is an example of a `for` loop used as a `while` loop:

```go
for condition {
    // code block to be executed
}
```

The `for` loop can also be used as a `do-while` loop by omitting the initialization statement and using a `break` statement to exit the loop. Here is an example of a `for` loop used as a `do-while` loop:

```go
for {
    // code block to be executed
    if condition {
        break
    }
}
```

### Example

Here is an example of a `for` loop that prints the numbers from 1 to 5:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
}
```

Output:

```
1 2 3 4 5
```

In this example, the `for` loop initializes the variable `i` to 1, checks if `i` is less than or equal to 5, and increments `i` by 1 after each iteration of the loop. The loop prints the value of `i` on each iteration until `i` is greater than 5.

### Infinite Loops

An infinite loop is a loop that never terminates. You can create an infinite loop by omitting the condition statement in the `for` loop. Here is an example of an infinite loop:

```go
for {
    fmt.Println("Hello, World!")
}
```

In this example, the `for` loop does not have a condition statement, so it will continue to execute the code block inside the loop indefinitely.

### Loop Control Statements

Go provides several loop control statements that allow you to control the flow of the loop. The loop control statements are:

- `break`: The `break` statement is used to exit the loop immediately.
- `continue`: The `continue` statement is used to skip the current iteration of the loop and continue with the next iteration.
- `goto`: The `goto` statement is used to jump to a specific label within the loop.

Here is an example of using the `break` statement to exit the loop when `i` is equal to 3:

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        break
    }
    fmt.Printf("%d ", i)
}
```

Output:

```
1 2
```

In this example, the loop prints the numbers from 1 to 5, but it exits the loop when `i` is equal to 3 using the `break` statement.

Here is an example of using the `continue` statement to skip the iteration when `i` is equal to 3:

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }
    fmt.Printf("%d ", i)
}
```

Output:

```go
1 2 4 5
```

In this example, the loop prints the numbers from 1 to 5, but it skips the iteration when `i` is equal to 3 using the `continue` statement.

Here is an example of using the `goto` statement to jump to a specific label within the loop:

```go
func main() {
    i := 1
loop:
    for i <= 5 {
        fmt.Printf("%d ", i)
        i++
        goto loop
    }
}
```

Output:

```
1 2 3 4 5
```

In this example, the loop prints the numbers from 1 to 5 using the `goto` statement to jump to the `loop` label within the loop.

### Range Loops

Go provides a `range` keyword that can be used to iterate over arrays, slices, strings, maps, and channels. The `range` keyword returns the index and value of each element in the collection.

Here is an example of using the `range` keyword to iterate over an array:

```go
func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
    for i, num := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", i, num)
    }
}
```

Output:

```
Index: 0, Value: 1
Index: 1, Value: 2
Index: 2, Value: 3
Index: 3, Value: 4
Index: 4, Value: 5
```
