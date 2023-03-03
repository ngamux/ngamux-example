package main

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()
	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).Text("GET /")
	})

	users := mux.Group("/users")
	users.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).Text("GET /users")
	})

	users.Get("/:id", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).Text(fmt.Sprintf("GET /users/:id with id = %s", ngamux.Req(r).Params("id")))
	})

	users.Get("/:id/:slug", func(rw http.ResponseWriter, r *http.Request) error {
		c := ngamux.NewCtx(rw, r)
		return c.Res.Json(map[string]string{
			"id":   c.Req.Params("id"),
			"slug": c.Req.Params("slug"),
		})
	})

	http.ListenAndServe(":8080", mux)
}
