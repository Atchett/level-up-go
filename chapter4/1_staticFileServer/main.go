package main

import (
	"log"
	"net/http"
)

func main() {

	//Simple static web server
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	log.Fatal(http.ListenAndServe(":3004", mux))

}
