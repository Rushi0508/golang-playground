package main

import (
	"errors"
	"fmt"
)

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a/b, nil;
}

func main() {
	// result, err := divide(10, 0)
	result, _ := divide(10,0);
    // if err != nil {
    //     fmt.Println("Error:", err)
    // } else {
    //     fmt.Println("Result:", result)
    // }
	fmt.Println(result)
}