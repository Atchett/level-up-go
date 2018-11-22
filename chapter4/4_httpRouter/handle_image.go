package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HandleImageNew = Diaplys new image form
func HandleImageNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic(fmt.Errorf("this should not be reached"))
}
