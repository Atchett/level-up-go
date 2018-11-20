package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HandleUserNew  - displays new user page
func HandleUserNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	RenderTemplate(w, r, "users/new", nil)
}
