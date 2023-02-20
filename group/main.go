package main

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()
	users := mux.Group("/users")
	users.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.String(rw, "GET /users")
	})

	http.ListenAndServe(":8080", mux)
}
