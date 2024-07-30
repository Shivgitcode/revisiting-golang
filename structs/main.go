package main

import "fmt"

func main() {
	fmt.Println("Struct in go lang")
	//no inheritence in golang; No super or parent
	shivnansh := User{"shivansh", "shivneeraj2004@gmail.com", true, 16}
	fmt.Println(shivnansh.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
