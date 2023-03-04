package main

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ctx"
	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()

	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).
			Status(http.StatusOK).
			Text("GET /")
	})

	mux.Get("/users", func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(rw, "GET /users")
		return nil
	})

	mux.Get("/products", func(rw http.ResponseWriter, r *http.Request) error {
		c := ctx.New(rw, r)
		return c.Res().
			Status(201).
			Text("GET /products")
	})

	http.ListenAndServe(":8080", mux)
}
