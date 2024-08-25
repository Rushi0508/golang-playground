## Error Handling and \_ Keyword

Go uses a simple error handling mechanism where functions return an `error` type along with the expected value. Additionally, the `_` keyword is used to ignore return values or variables that are not needed.

### Error handling

In Go, errors are returned as the last value from a function, and the caller is expected to check for the error.

For error we use specific type called `error`

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### Underscore `_` Keyword

The underscore (\_) keyword is used to ignore values that a function returns but which are not needed.

Mostly `_` are used to ignore the error returned by the function so that we don't need to handle them.

Example 1:

```go
func divide(a, b float32)(float32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a/b, nil;
}

func main() {
	result, _ := divide(10,0);
	fmt.Println(result)
}
```

Example 2:

```go
func calculate(a, b int) (int, int) {
    return a + b, a * b
}

func main() {
    sum, _ := calculate(3, 4)
    fmt.Println("Sum:", sum)
}
```

In the above example, the second return value (the product) is ignored using the `_` keyword. Only the sum is captured.
