package main

import "fmt"

func main() {
	// Arrays
	var names [5]string
	names[0] = "Rushi"
	names[1] = "Jack"
	fmt.Println("Names are: ", names)

	var nums = [10]int{1,2,3,4,5}
	fmt.Println("Nums are: ", nums)

	fmt.Println("Length of nums is: ", len(nums))
	fmt.Println("Element at index 3 is: ", nums[3])

	// Slices
	var fruits = []string{"Apple", "Banana", "Orange"}
	fmt.Println("Fruits are: ", fruits)

	fmt.Println("Length of fruits is: ", len(fruits))
	fmt.Println("Element at index 1 is: ", fruits[1])
	fmt.Println("Capacity of fruits is: ", cap(fruits))

	var numbers = make([]int, 3, 5)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	fmt.Println("Slice is: ", numbers)
	fmt.Println("Length of slice is: ", len(numbers))
	fmt.Println("Capacity of slice is: ", cap(numbers))

	numbers = append(numbers, 4)
	fmt.Println("Slice is: ", numbers)
	fmt.Println("Length of slice is: ", len(numbers))
	fmt.Println("Capacity of slice is: ", cap(numbers))


}