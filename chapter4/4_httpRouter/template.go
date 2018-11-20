package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var layoutFuncs = template.FuncMap{
	"yeild": func() (string, error) {
		return "", fmt.Errorf("yeild called inappropriately")
	},
}

var layout = template.Must(
	template.
		New("layout.gohtml").
		Funcs(layoutFuncs).
		ParseFiles("templates/layout.gohtml"),
)

var templates = template.Must(
	template.
		New("t").
		ParseGlob("templates/**/*.gohtml"),
)

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

	funcs := template.FuncMap{
		"yeild": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	}

	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)

	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}
