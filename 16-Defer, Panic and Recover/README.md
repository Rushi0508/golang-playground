## Defer, Panic and Recover in Go

### Defer Keyword

The defer keyword is used to defer the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

The defer keyword is often used to close files, unlock mutexes, and to do cleanup tasks.

```go
package main

import "fmt"

func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

Output

```
Hello
World
```

In the above example, the `fmt.Println("World")` is deferred until the `main` function returns.

#### Multiple Deferred Calls

You can have multiple deferred calls in a function. The deferred calls are executed in the reverse order of their appearance.

```go
func main() {
    defer fmt.Println("World")
    defer fmt.Println("Hello")
}
```

Output

```
Hello
World
```

In the above example, the `fmt.Println("Hello")` is deferred first, and the `fmt.Println("World")` is deferred next. The deferred calls are executed in the reverse order of their appearance.

```go
func main(){
    fmt.Println("Start")
    defer fmt.Println("Middle")
    defer fnt.Println("Sum is " , add(1,2))
    fmt.Println("End")
}

func add(a,b int) int{
    return a + b
}
```

Output

```go
Start
End
Sum is 3
Middle
```

#### Defer and Return Values

The deferred function can access and modify the return values of the surrounding function.

```go
func main() {
    fmt.Println(add(1, 2))
}

func add(a, b int) int {
    defer fmt.Println("Calculating the sum")
    return a + b
}
```

Output

```
Calculating the sum
3
```

In the above example, the `fmt.Println("Calculating the sum")` is deferred until the `add` function returns. The deferred function can access and modify the return values of the surrounding function.

### Panic

The panic function is used to terminate the program abruptly. When a function calls panic, the program stops executing the current function and all the deferred functions are executed. The panic function is similar to the exceptions in other programming languages.

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")
    panic("Something went wrong")
    fmt.Println("End")
}
```

Output

```
Start
panic: Something went wrong

goroutine 1 [running]:
main.main()
        /tmp/sandbox282013013/main.go:7 +0x6b
```

```go
fmt.Println("start")
defer fmt.Println("middle")
panic("something happened")
fmt.Println("end")
```

Output

```
start
middle
panic: something happened

goroutine 1 [running]:
main.main()
        /tmp/sandbox282013013/main.go:7 +0x6b
```

In the above example, the `fmt.Println("end")` is not executed because the program panics before that.

### Recover

The recover function is used to regain control of a panicking goroutine. The recover function is only useful when called inside a deferred function. The recover function stops the panic and returns the value that was passed to the call to panic.

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("Something went wrong")
    fmt.Println("End")
}
```

Output

```
Start
Recovered from panic: Something went wrong

goroutine 1 [running]:
main.main()
        /tmp/sandbox282013013/main.go:7 +0x6b
```

In the above example, the `fmt.Println("End")` is not executed because the program panics before that. The `recover` function regains control of the panicking goroutine.

### Best Practices

- The panic function should be used only for unrecoverable errors. The panic function is similar to exceptions in other programming languages.

- The recover function should be used to regain control of a panicking goroutine. The recover function is only useful when called inside a deferred function.

- The defer keyword is often used to close files, unlock mutexes, and to do cleanup tasks.

- The deferred calls are executed in the reverse order of their appearance.

- The deferred function can access and modify the return values of the surrounding function.

- The panic function stops the program abruptly.

- The recover function stops the panic and returns the value that was passed to the call to panic.

### Conclusion

The defer keyword is used to defer the execution of a function until the surrounding function returns. The panic function is used to terminate the program abruptly. The recover function is used to regain control of a panicking goroutine.
