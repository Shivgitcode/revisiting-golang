package main

import "fmt"

func main() {
	branch := make(map[string]string)

	branch["IT"] = "Information Technology"
	branch["B.Tech"] = "Bachelor in Technology"

	fmt.Println(branch["IT"])

	for key, value := range branch {
		fmt.Printf("%v : %v\n", key, value)
	}

}
