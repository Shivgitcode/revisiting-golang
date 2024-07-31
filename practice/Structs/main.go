package main

import "fmt"

func main() {

	user1 := User{
		Name:   "shivansh",
		rollno: 2210990830,
		group:  "G-12",
	}

	fmt.Println(user1.Name)

}

type User struct {
	Name   string
	rollno int
	group  string
}
