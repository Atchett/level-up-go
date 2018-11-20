package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// not in a function will run immediately
// creating a templates instance "t"
// parse files match the glob "templates/**/*.html"
// template.Must returns the template if valid, if not it will error
// means fail to start rather than fail on first request
var templates = template.Must(template.New("t").ParseGlob("templates/**/*.gohtml"))

var errorTemplate = `
	<html>
		<body>
			<h1>Error rendering template %s</h1>
			<p>%s</p>
		</body>
	</html>
`

// RenderTemplate is used to render templates
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}
