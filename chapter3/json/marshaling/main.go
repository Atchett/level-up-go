package main

import (
	"encoding/json"
	"fmt"
)

//Article stores some simeple data for the example
type Article struct {
	Name       string
	AuthorName string
	draft      bool
}

func main() {

	article := Article{
		Name:       "JSON in Go",
		AuthorName: "John Spurgin",
		draft:      false,
	}

	data, err := json.Marshal(article)

	if err != nil {
		fmt.Println("Could not marshal article", err)
	} else {
		fmt.Println(string(data))
	}

}
