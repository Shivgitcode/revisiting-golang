package main

import "fmt"

func main() {
	user1 := User{"shivansh", "shivneeraj2004@gmail.com", true, 20}
	user1.GetStatus()
	user1.NewMail()
	fmt.Println(user1.email)

}

type User struct {
	name   string
	email  string
	status bool
	age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.status)

}

func (u User) NewMail() {
	u.email = "hello@gmail.com"
	fmt.Println("Email of this user is : ", u.email)
}
