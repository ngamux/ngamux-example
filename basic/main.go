package main

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()

	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).
			String("GET /")
	})

	mux.Get("/users", func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(rw, "GET /users")
		return nil
	})

	mux.Get("/products", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).
			Status(201).
			String("GET /products")
	})

	http.ListenAndServe(":8080", mux)
}
