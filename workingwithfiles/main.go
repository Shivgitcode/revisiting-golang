package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./myTextFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./myTextFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is ", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
