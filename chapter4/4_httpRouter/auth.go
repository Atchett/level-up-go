package main

import (
	"net/http"
)

// AuthenticateRequest - redirects the user if they are not authenticated
func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	//Redirect the user to the login if they are not authenticated
	authenticated := false
	if !authenticated {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}
