package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HandleHome - Displays the home page
func HandleHome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	RenderTemplate(w, r, "index/home", nil)
}
