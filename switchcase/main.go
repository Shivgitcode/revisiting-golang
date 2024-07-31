package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	fmt.Println("Switch and case in golang")
	// rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
		fallthrough
	default:
		fmt.Println("You can move spot ", diceNumber)
	}

	// fmt.Println("value of dice is", diceNumber)

}
