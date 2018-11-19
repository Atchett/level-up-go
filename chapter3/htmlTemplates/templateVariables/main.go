package main

import (
	"html/template"
	"os"
)

//Product example used for code
type Product struct {
	Price    float64
	Quantity float64
}

// Multiply accepts 2 float64 values and returns a float64 value
func Multiply(a, b float64) float64 {
	return a * b
}

func main() {

	tmpl := template.New("Foo")
	tmpl.Funcs(template.FuncMap{"multiply": Multiply})

	tmpl, err := tmpl.Parse(`
		{{$total := multiply .Price .Quantity}}
		Price: ${{printf "%.2f" $total }}
	`)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, Product{
		Price:    12.3,
		Quantity: 2,
	})

	if err != nil {
		panic(err)
	}

}
