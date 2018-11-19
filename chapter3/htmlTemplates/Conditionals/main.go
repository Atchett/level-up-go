package main

import (
	"html/template"
	"os"
)

//Article stores some simeple data for the example
type Article struct {
	Name       string
	AuthorName string
	Draft      bool
}

func main() {
	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
		Draft:      true,
	}

	tmpl, err := template.New("Foo").Parse(
		"{{.Name}}{{if .Draft}} (Draft) {{end}}\n",
	)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}
}
