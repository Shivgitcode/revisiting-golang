package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a index: ")
	input1, _ := reader.ReadString('\n')
	index, _ := strconv.ParseInt(strings.TrimSpace(input1), 36, 64)
	fmt.Println(index)
	arr1 = append(arr1[0:index], arr1[index+1:]...)
	fmt.Println(arr1)

}
