package main

import "fmt"

// declaring structs
type Person struct{
	Name string
	Age int
}

// struct embedding
type Employee struct{
	Name string
	Age int
	contact_details Contact
}

type Contact struct{
	Email string
	Phone string
}

func getAge(p Person) int{
	return p.Age;
}

func main(){
	// Creating structs
	var person1 Person
	person1.Name = "Rushi"
	person1.Age = 21
	fmt.Println("Using method 1")
	fmt.Println(person1)

	person2 := Person{
		Name: "Rushi",
		Age: 21,
	}
	fmt.Println("Using method 2")
	fmt.Println(person2)

	person3 := new(Person)
	person3.Name = "Rushi"
	person3.Age = 21
	fmt.Println("Using method 2")
	fmt.Println(person3)

	// structs using functions
	var age int = getAge(person1);
	fmt.Println("Age is ", age)

	// struct embedding
	contact := Contact{
		Email: "hello@gmail.com",
		Phone: "1234567890",
	}
	employee := Employee{
		Name: "Rushi",
		Age: 21,
		contact_details: contact,
	}
	fmt.Println(employee)
}

