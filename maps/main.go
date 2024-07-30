package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	animes := []string{"one piece", "black clover", "mha"}

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

	for key, value := range animes {
		fmt.Println(key+1, "anime: ", value)
	}

}
