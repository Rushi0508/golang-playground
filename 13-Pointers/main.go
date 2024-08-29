package main

func main() {
	// Declaring a pointer
	var ptr *int
	println(ptr) // <nil>

	// Declaring a pointer with a value
	var a int = 10
	ptr = &a
	println(ptr) // 0xc0000b6010

	// Dereferencing a pointer
	println(*ptr) // 10

	// Changing the value of a variable through a pointer
	*ptr = 20
	println(a) // 20
}