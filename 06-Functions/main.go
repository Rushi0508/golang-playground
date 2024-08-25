package main

import "fmt"

// function without parameters and return type
func sayHello() {
	fmt.Println("Hello World")
}

// function with parameters and return type
func getSum(a,b int)int{
	return a+b;
}

// function with named return type
func getProduct(a,b int) (result int){
	result = a*b;
	return
}

// function with multiple return values
func sumAndDifference(a,b int)(int, int){
	return a+b,a-b;
}

// function with multiple named return
func sumAndProduct(a,b int)(sum int, pro int){
	sum = a+b;
	pro = a*b;
	return 
}
func main() {
	sayHello()
	
	sum := getSum(1,2)
	fmt.Println(sum)

	pro := getProduct(1,2)
	fmt.Println(pro)

	addi,subs := sumAndDifference(1,2)
	fmt.Println(addi , subs)

	s,p := sumAndProduct(1,2)
	fmt.Println(s , p)

}