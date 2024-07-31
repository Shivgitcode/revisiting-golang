package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Printf("go")
	// PerformGetRequest()
	PerformPostRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:2000/api/v1/todo"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(response.Body)

	var responseString strings.Builder

	fmt.Println(string(data))
	byteCount, _ := responseString.Write(data)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	defer response.Body.Close()

	fmt.Println("Status code is", response.StatusCode)

}

func PerformPostRequest() {
	const url = "http://localhost:2000/api/v1/todo"

	responseBody := strings.NewReader(`
	{
		"todoName":"make notes",
		"isDone":false
	}
	`)

	response, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var responseString strings.Builder
	data, _ := ioutil.ReadAll(response.Body)
	byteCode, _ := responseString.Write(data)
	fmt.Println(byteCode)
	fmt.Println(responseString.String())

}
