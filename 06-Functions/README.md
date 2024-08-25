## Functions

- Functions are essential part of Go.
- Functions are declared using `func` keyword, followed by function name, parameters(if any), return type(if any), and the function body.
- As we know `main` is the entry point of the program.
- The function body is exclosed in curly braces. In Go, the opening curly brace `{` should be on the same line as function declaration.

```go
// Correct
func hello(){

}

// Incorrect
func hello()
{

}
```

### Function without parameters and return type

```go
func sayHello() {
	fmt.Println("Hello World")
}
```

### Function with parameters and return type

```go
func getSum(a,b int) int {
	return a+b;
}
```

We must specify the type for each of the parameter. But if all the parameters have the same type we can write the type at last.

### Function with named return type

```go
func getProduct(a,b int) (result int){
	result = a*b;
	return
}
```

### Function with multiple return types

```go
func sumAndDifference(a,b int)(int, int){
	return a+b,a-b;
}
```

### Function with multiple named return types

```go
func sumAndProduct(a,b int)(sum int, pro int){
	sum = a+b;
	pro = a*b;
	return
}
```
