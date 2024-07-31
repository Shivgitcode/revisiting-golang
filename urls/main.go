package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("Welcome to handling URLs in golang")

	response, err := http.Get(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)

	defer response.Body.Close()

}
