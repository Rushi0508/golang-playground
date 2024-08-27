## Maps in Go

A map is a collection of key-value pairs.

Each key in a map must be unique, and the key-value pairs are unordered.

Maps are used to store and retrieve values based on a key. In Go, maps are created using the `make` function or a map literal.

### Creating a Map

To create a map, you can use the `make` function or a map literal.

Here is an example of creating a map using the `make` function:

```go
func main() {
    // Create a map using the make function
    // <name> : <runs>
    scoreBoard := make(map[string]int)

    // Add key-value pairs to the map
    scoreBoard["Alice"] = 42
    scoreBoard["Bob"] = 34
    scoreBoard["Charlie"] = 56

    // Print the map
    fmt.Println(scoreBoard)
}
```

Output:

```
map[Alice:42 Bob:34 Charlie:56]
```

In this example, we create a map called `scoreBoard` using the `make` function and add key-value pairs to the map.

Here is an example of creating a map using a map literal:

```go
func main() {
    // Create a map using a map literal
    scoreBoard := map[string]int{
        "Alice":   42,
        "Bob":     34,
        "Charlie": 56,
    }

    // Print the map
    fmt.Println(scoreBoard)
}
```

Output:

```
map[Alice:42 Bob:34 Charlie:56]
```

In this example, we create a map called `scoreBoard` using a map literal.

### Accessing Elements in a Map

You can access elements in a map using the key.

Here is an example of accessing elements in a map:

```go
// Access elements in the map
fmt.Println(scoreBoard["Alice"])
// Output: 42

fmt.Println(scoreBoard["Bob"])
// Output: 34

fmt.Println(scoreBoard["Charlie"])
// Output: 56
```

In this example, we access elements in the `scoreBoard` map using the keys `"Alice"`, `"Bob"`, and `"Charlie"`.

### Updating Elements in a Map

You can update elements in a map by assigning a new value to the key.

Here is an example of updating elements in a map:

```go
// Update elements in the map
scoreBoard["Alice"] = 50
scoreBoard["Bob"] = 40
scoreBoard["Charlie"] = 60

// Print the updated map
fmt.Println(scoreBoard)
```

Output:

```
map[Alice:50 Bob:40 Charlie:60]
```

In this example, we update the elements in the `scoreBoard` map by assigning new values to the keys `"Alice"`, `"Bob"`, and `"Charlie"`.

### Deleting Elements from a Map

You can delete elements from a map using the `delete` function.

Here is an example of deleting elements from a map:

```go
// Delete elements from the map
delete(scoreBoard, "Bob")

// Print the updated map
fmt.Println(scoreBoard)
```

Output:

```
map[Alice:50 Charlie:60]
```

In this example, we delete the element with the key `"Bob"` from the `scoreBoard` map.

### Checking if a Key Exists in a Map

Here is an example of checking if a key exists in a map:

```go
value, ok := scoreBoard["Alice"]
if ok {
    fmt.Println("Alice's score is", value)
} else {
    fmt.Println("Alice's score is not available")
}
```

If the key `"Alice"` exists in the `scoreBoard` map, the variable `value` will contain the value associated with the key, and the variable `ok` will be `true`.

If the key does not exist, the variable `value` will contain the zero value of the map's value type, and the variable `ok` will be `false`.

Output:

```
Alice's score is 50
```

In this example, we check if the key `"Alice"` exists in the `scoreBoard` map and print the score if it exists.

### Iterating Over a Map

You can iterate over a map using a `for` loop.

Here is an example of iterating over a map:

```go
// Iterate over the map
for key, value := range scoreBoard {
    fmt.Println(key, ":", value)
}
```

Output:

```
Alice : 50
Charlie : 60
```

In this example, we iterate over the `scoreBoard` map and print the key-value pairs.
