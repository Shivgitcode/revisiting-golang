package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", &ptr)
	myNumber := 23

	fmt.Println("Address of myNumber: ", &myNumber)

	var ptr *int = &myNumber
	// fmt.Println("printing ptr : ", ptr)
	fmt.Println(myNumber)
	*ptr = *ptr + 2
	fmt.Println(myNumber)

}
