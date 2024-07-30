package main

import "fmt"

func main() {
	fmt.Println("variables")

	var username string = "shivansh"
	var isLoggedIn bool = false
	var smallFloat float32 = 255
	// var smallVal int64 = 45
	// fmt.Println(username)
	fmt.Printf("Variable is of type: %s\n", username)
	fmt.Printf("isLoggedIn %T", isLoggedIn)
	fmt.Printf("float value %0.2f", smallFloat)

	//implicit type;

	var website = "hello"
	fmt.Println(website)
	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

}
