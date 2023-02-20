package main

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()
	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.String(rw, "GET /")
	})

	users := mux.Group("/users")
	users.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.String(rw, "GET /users")
	})

	users.Get("/:id", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.String(rw, fmt.Sprintf("GET /users/:id with id = %s", ngamux.GetParam(r, "id")))
	})

	users.Get("/:id/:slug", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.JSON(rw, map[string]string{
			"id":   ngamux.GetParam(r, "id"),
			"slug": ngamux.GetParam(r, "slug"),
		})
	})

	http.ListenAndServe(":8080", mux)
}
