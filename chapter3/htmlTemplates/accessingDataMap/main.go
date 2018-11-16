package main

import (
	"html/template"
	"os"
)

func main() {
	article := map[string]string{
		"Name":       "The Go html/template package",
		"AuthorName": "Mal Curtis",
	}

	tmpl, err := template.New("Foo").Parse("'{{.Name}}' by {{.AuthorName}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, article)
	if err != nil {
		panic(err)
	}
}
