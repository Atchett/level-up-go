package main

import (
	"html/template"
	"os"
)

func main() {

	//formatting
	tmpl, _ := template.New("Foo").Parse(
		"Price: ${{printf \"%.2f\" .}}\n",
	)
	tmpl.Execute(os.Stdout, 12.3)

	// alternate with pipeline
	tmpl2, _ := template.New("Foo").Parse(
		"Price: ${{. | printf \"%.2f\"}}\n",
	)
	tmpl2.Execute(os.Stdout, 12.3)

}
