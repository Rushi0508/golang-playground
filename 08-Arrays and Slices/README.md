## Arrays & Slices in Go

### Arrays

Arrays are the same that we use in the other programming languages. It's a collection of elements of the same type. The size of the array is fixed and cannot be changed once it's declared.

#### Array Declaration

```go
var names [5]string
names[0] = "Rushi"
names[1] = "Jack"
fmt.Println("Names are: ", names)
```

#### Array Declaration with initialization

```go
var nums = [10]int{1,2,3,4,5}
fmt.Println("Nums are: ", nums)
```

In Go, when you declare an array or [slice](./README.md#slices), the elements are initialized to the default value of array type. For numeric types(int, float etc) it's initialised to `0`, for string it's empty string `""`, for boolean it's `false` etc.

#### Length of array

```go
len(nums)
```

#### Access array element

```go
fmt.Println("Element at 2nd index is: ", nums[2])
```

### Slices

In Go, slices are more powerful, flexible and convenient than arrays. Slices are like dynamic arrays with the ability to resize itself. Slices are more common than arrays in Go.

#### Slice Declaration

```go
var names = []string{}
```

#### Slice Declaration with initialization

```go
var names = []string{"Rushi", "Jack", "John"}
fmt.Println("Names are: ", names)
```

#### Length of slice

```go
len(names)
```

#### Capacity of slice

```go
cap(names)
```

Here the capacity of the slice is the maximum number of elements that the slice can hold. When the length of the slice reaches the capacity, the slice will be resized to double the capacity.

#### Initialize slice with make

```go
var names = make([]string, 5, 10)
fmt.Println("Names are: ", names)
```

- In the above example, we have created a slice with a length of 5 and a capacity of 10.
- First argument is the type of slice, second argument is the length and third argument is the capacity.
- Where the length is the number of elements in the slice and the capacity is the maximum number of elements that the slice can hold.

#### Access slice element

```go
fmt.Println("Element at 2nd index is: ", names[2])
```

#### Append to slice

```go
names = append(names, "Doe")
fmt.Println("Names are: ", names)
```

In the above example, we have appended a new element to the slice. The `append` function returns a new slice with the new element added at the end of the slice.

#### Slice of slice

```go
var newNames = names[1:3]
fmt.Println("New names are: ", newNames)
```

In the above example, we have created a new slice from the existing slice. The new slice will contain elements from index 1 to 3(exclusive) of the existing slice.
