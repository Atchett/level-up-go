package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/articles/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from /articles/latest")
	})

	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from /articles/")
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from /users")
	})

	mux.HandleFunc("/redirected", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from /redirected")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Something went wrong", 500)
	})

	mux.HandleFunc("/s", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:3001/redirected", 301)
	})

	http.ListenAndServe(":3001", mux)

}
