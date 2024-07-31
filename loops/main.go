package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	animes := []string{"One Piece", "My Hero academia", "Black Clover", "Dragon Ball Z", "Blue Lock"}

	for key, value := range animes {
		fmt.Printf("%v : %v\n", key, value)

	}

	for d := 0; d < len(animes); d++ {
		fmt.Println(animes[d])

	}

	num := 1

	for num < 10 {
		fmt.Println("value is:", num)
		if num == 5 {
			num++
			continue
		}
		num++
	}
}
