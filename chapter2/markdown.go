package main

import (
	"fmt"

	"github.com/russross/blackfriday"
)

func main() {
	// note : NO indentation on the text
	markdown := []byte(`
#this is a header
* and
* this
* is
* a
* list
	`)

	html := blackfriday.MarkdownBasic(markdown)
	fmt.Println(string(html))
}
