package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"courseName"`
	Price    int    `json:"price"`
	Platform string
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []Course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename":"ReactJS Bootcamp",
			"Price":299,
			"website":"LearnCodeOnline.in",
			"tags":["web-dev","js"]
		}
	`)

	var lcoCourse Course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
}
