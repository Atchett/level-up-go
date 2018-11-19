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

	// create a slice of Articles
	var s []Article

	// create some articles and append them to the slice
	a1 := Article{
		Name:       "Some Name",
		AuthorName: "Bob",
		Draft:      false,
	}
	s = append(s, a1)

	a2 := Article{
		Name:       "Name 2",
		AuthorName: "Frank",
		Draft:      true,
	}
	s = append(s, a2)

	tmpl, err := template.New("Foo").Parse(`

		{{define "ArticleResource"}}
			<p>{{.Name}} by {{.AuthorName}}</p>
		{{end}}

		{{define "ArticleLoop"}}
			{{range .}}
				{{template "ArticleResource" .}}
			{{else}}
				<p>No published articles yet</p>
			{{end}}
		{{end}}

		{{template "ArticleLoop" .}}
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}
}
