## Structs in Go

Structs are a way to create complex data types in Go. They are similar to classes in object-oriented programming languages. Structs are used to group data together to create records.

### Declaring a Struct

To declare a struct, you use the `type` keyword followed by the name of the struct and the keyword `struct`. You then define the fields of the struct inside curly braces.

```go
type Person struct {
    Name string
    Age int
}
```

### Creating a Struct

To create a struct, you use the `var` keyword followed by the name of the variable, the name of the struct, and the values for the fields.

- 1st method

```go
var person Person

person.Name = "Alice"
person.Age = 30

fmt.Println(person);
```

- 2nd Method

```go
person := Person{
    Name: "Alice",
    Age: 30
}
fmt.Println(person)
```

- 3rd Method (using new keyword)

```go
person:= new(Person)
person.Name = "Alice"
person.Age = 30

fmt.Println(person)
fmt.Println("Age is: ", person.Age);
```

### Struct Methods

You can define methods on structs in Go. To do this, you define a function with a receiver. The receiver is the name of the struct type you want to define the method on.

```go
func getAge(p Person) int{
	return p.Age;
}
```

### Struct Embedding

Go supports a feature called struct embedding, which allows you to include one struct as a field in another struct.

```go
type Address struct {
    City string
    State string
}

type Person struct {
    Name string
    Age int
    Address: Address
}

func main() {
    person := Person{
        Name: "Alice",
        Age: 30,
        Address: Address{
            City: "New York",
            State: "NY",
        },
    }

    fmt.Println(person)
}
```
