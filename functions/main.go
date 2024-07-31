package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter("shivansh")
	greeter("Shweta")
	num := sum(10, 30)
	num2, message := proAdder(10, 30, 40, 50)
	fmt.Println(num)
	fmt.Println(num2, message)
}

func greeter(name string) {
	fmt.Printf("Hello %v\n", name)
}
func sum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func proAdder(values ...int) (int, string) {
	total := 0
	for i := 0; i < len(values); i++ {
		total = total + values[i]
	}
	return total, "these are total values"
}
