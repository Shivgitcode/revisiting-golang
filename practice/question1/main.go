// taiking input from user and printing sum

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
	fmt.Println("Enter num 1: ")
	input1, _ := reader.ReadString('\n')
	num1, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	fmt.Println("Enter num 2: ")
	input2, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	sum := num1 + num2
	fmt.Println("Sum of num1 and num2 is : ", sum)

}
