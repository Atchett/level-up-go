package main

import (
	"fmt"
	"net/http"
	"time"
)

// UptimeHandler writes the number of seconds since
// starting the response
type UptimeHandler struct {
	Started time.Time
}

//NewUptimeHandler - Retuens the time since started
func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w,
		fmt.Sprintf("Current uptime: %s", time.Since(h.Started)),
	)
}

//SecretTokenHandler secures a request with a secret token
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

//ServeHTTP makes SecretTokenHandler implement the http.Handler interface
func (h SecretTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// checks the query string for the secret token
	if r.URL.Query().Get("secret_token") == h.secret {
		// the secret token matched, call the next handler
		h.next.ServeHTTP(w, r)
	} else {
		// No match, return a 404 Not Found response
		http.NotFound(w, r)
	}

}

func main() {
	http.Handle("/", SecretTokenHandler{
		next:   NewUptimeHandler(),
		secret: "SomeStuff",
	})
	http.ListenAndServe(":3003", nil)
}
