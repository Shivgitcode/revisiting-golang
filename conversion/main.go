package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your name:")
	// welcome, _ := reader.ReadString('\n')
	fmt.Println("Enter rating:")
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

	// fmt.Println("Hello", welcome)

}
