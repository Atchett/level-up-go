package main

import (
	"encoding/json"
	"fmt"
)

// Product is an example struct
type Product struct {
	//Field appears in JSON with the key "name"
	Name string `json:"name"`
	// Field appears in the JSON with the key "author_name",
	// but doesn't appear at all if its value is empty
	AuthorName string `json:"author_name,omitempty"`

	//Field will not appear in the JSON representation
	CommissionPrice float64 `json:"-"`
}

func main() {

	product := Product{
		Name:            "Shampoo",
		AuthorName:      "Weird field for Shampoo",
		CommissionPrice: 2.35,
	}

	data, err := json.MarshalIndent(product, "", "    ")

	if err != nil {
		fmt.Println("Not able to marshal", err)
	} else {
		fmt.Println(string(data))
	}
}
