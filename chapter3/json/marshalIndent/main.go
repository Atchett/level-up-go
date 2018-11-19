package main

import (
	"encoding/json"
	"fmt"
)

//Article stores some simeple data for the example
type Article struct {
	Name       string
	AuthorName string
	Draft      bool
}

func main() {

	article := Article{
		Name:       "JSON in Go",
		AuthorName: "John Spurgin",
		Draft:      false,
	}

	data, err := json.MarshalIndent(article, "", "    ")

	if err != nil {
		fmt.Println("Could not marshal article", err)
	} else {
		fmt.Println(string(data))
	}

}
