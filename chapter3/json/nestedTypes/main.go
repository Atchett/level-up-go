package main

import (
	"encoding/json"
	"fmt"
)

// Article is an example struct
type Article struct {
	Name string
}

// ArticleCollection is an example struct
type ArticleCollection struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
}

func main() {
	p1 := Article{
		Name: "JSON in Go",
	}
	p2 := Article{
		Name: "Marshaling is good",
	}
	articles := []Article{p1, p2}
	collection := ArticleCollection{
		Articles: articles,
		Total:    len(articles),
	}

	data, err := json.MarshalIndent(collection, "", "    ")
	if err != nil {
		fmt.Println("Cannot marshal", err)
	} else {
		fmt.Println(string(data))
	}
}
