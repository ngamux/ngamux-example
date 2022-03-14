package main

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()

	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(rw, "GET /")
		return nil
	})

	mux.Get("/users", func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(rw, "GET /users")
		return nil
	})

	http.ListenAndServe(":8080", mux)
}
