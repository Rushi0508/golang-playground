package main

import "fmt"

func main() {
	var names [5]string
	names[0] = "Rushi"
	names[1] = "Jack"
	fmt.Println("Names are: ", names)

	var nums = [10]int{1,2,3,4,5}
	fmt.Println("Nums are: ", nums)

	fmt.Println("Length of nums is: ", len(nums))
}