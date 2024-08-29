## Strings Package in Go

As we know google developed Go because the entire web scraping is based on strings so there is lot of data manipulation on texts so Go has a very powerful `strings` package.

### Strings Package Functions

1. **Contains**: It checks if a string contains a substring or not.

```go
data := "Hello, World!"
fmt.Println(strings.Contains(data, "Hello")) // true
```

2. **Count**: It counts the number of non-overlapping instances of a substring in a string.

```go
data := "Hello, World!"
fmt.Println(strings.Count(data, "l")) // 3
```

3. **HasPrefix**: It checks if a string starts with a prefix or not.

```go
data := "Hello, World!"
fmt.Println(strings.HasPrefix(data, "Hello")) // true
```

4. **HasSuffix**: It checks if a string ends with a suffix or not.

```go
data := "Hello, World!"
fmt.Println(strings.HasSuffix(data, "World!")) // true
```

5. **Index**: It returns the index of the first instance of a substring in a string.

```go
data := "Hello, World!"
fmt.Println(strings.Index(data, "World")) // 7
```

6. **Join**: It concatenates the elements of a slice of strings to create a single string.

```go
data := []string{"Hello", "World"}
fmt.Println(strings.Join(data, ", ")) // Hello, World
```

7. **Repeat**: It returns a new string consisting of count copies of the string.

```go
data := "Hello, World!"
fmt.Println(strings.Repeat(data, 2)) // Hello, World!Hello, World!
```

8. **Replace**: It returns a copy of the string with the first n non-overlapping instances of old replaced by new.

```go
data := "Hello, World!"
fmt.Println(strings.Replace(data, "World", "Go", 1)) // Hello, Go!
```

here 1 is the number of replacements.

9. **Split**: It splits the string into all substrings separated by separator and returns a slice of the substrings.

```go
data := "Hello, World!"
fmt.Println(strings.Split(data, ", ")) // [Hello World!]
```

10. **ToLower**: It returns a copy of the string with all Unicode letters mapped to their lower case.

```go
data := "Hello, World!"
fmt.Println(strings.ToLower(data)) // hello, world!
```

11. **ToUpper**: It returns a copy of the string with all Unicode letters mapped to their upper case.

```go
data := "Hello, World!"
fmt.Println(strings.ToUpper(data))
```

12. **Trim**: It returns a slice of the string with all leading and trailing Unicode code points contained in cutset removed.

```go
data := "Hello, World!"
fmt.Println(strings.Trim(data, "H!")) // ello, World
```

And many more...
