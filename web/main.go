package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {

	fmt.Println("web request")
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("inside error")
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)

}
