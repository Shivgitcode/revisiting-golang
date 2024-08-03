package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url string = "http://localhost:2000/api/v1/todo"

type Todo struct {
	todoName string
	isDone   bool
}

func main() {
	fmt.Println("Making a get post request")
	GetRequest()
	// PostRequest()

}

func GetRequest() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	res, _ := ioutil.ReadAll(response.Body)

	var finalJson map[string]interface{}

	checKvalidJson := json.Valid(res)
	if checKvalidJson {
		fmt.Println("json was valid")
		json.Unmarshal(res, &finalJson)
		// fmt.Println(finalJson)

	} else {
		fmt.Println("json not valid")
	}

	delete(finalJson, "message")

	// var allTodos []Todo
	allTodos := finalJson["allTodos"]
	fmt.Println(allTodos)

}

func PostRequest() {
	responseBody := strings.NewReader(`
	{
		"todoName":"going to class",
		"isDone":true
	}

	`)

	response, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		panic(err)

	}
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(res, "", "\t")
	defer response.Body.Close()

	fmt.Println(jsonData)

}
