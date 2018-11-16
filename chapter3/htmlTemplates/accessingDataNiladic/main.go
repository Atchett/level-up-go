package main

import (
	"fmt"
	"html/template"
	"os"
)

//Article stores some simeple data for the example
type Article struct {
	Name       string
	AuthorName string
}

// ByLine writes the byline
func (a Article) ByLine() string {
	return fmt.Sprintf("Written by %s", a.AuthorName)
}

func main() {
	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
	}

	tmpl, err := template.New("Foo").Parse("{{.ByLine}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}
}
