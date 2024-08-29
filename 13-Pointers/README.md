## Pointers in Go

Pointers are a powerful feature of the Go programming language. They are used to store the memory address of another variable. They are used to store the address of a variable in order to access it.

They provide a way to directly work with memory.

### Declaring Pointers

A pointer is declared using the `*` symbol followed by the type of the variable it is pointing to.

```go
var ptr *int
```

### Initializing Pointers

A pointer can be initialized using the address of the variable it is pointing to.

```go
var a int = 10
var ptr *int = &a
```

### Dereferencing Pointers

Dereferencing a pointer means accessing the value of the variable it is pointing to. It is done using the `*` symbol.

```go
var a int = 10
var ptr *int = &a

fmt.Println(*ptr) // Output: 10
```

### Zero Value of a Pointer

The zero value of a pointer is `nil`.

```go
var ptr *int
fmt.Println(ptr) // Output: <nil>
```

### Pointer to Pointer

A pointer to a pointer is a pointer that points to another pointer.

```go
var a int = 10
var ptr *int = &a
var ptrPtr **int = &ptr

fmt.Println(**ptrPtr) // Output: 10
```

### Changing the Value of a Variable Using a Pointer

A pointer can be used to change the value of a variable.

```go
var a int = 10
var ptr *int = &a

*ptr = 20

fmt.Println(a) // Output: 20
```

### Passing Pointers to Functions

Pointers can be passed to functions as arguments.

```go
func changeValue(ptr *int) {
    *ptr = 20
}

var a int = 10
var ptr *int = &a

changeValue(ptr)

fmt.Println(a) // Output: 20
```
