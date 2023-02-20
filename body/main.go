package main

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()
	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.String(rw, "GET /")
	})

	mux.Post("/", func(rw http.ResponseWriter, r *http.Request) error {
		in := map[string]string{}
		err := ngamux.GetJSON(r, &in)
		if err != nil {
			ngamux.JSONWithStatus(rw, http.StatusBadRequest, err.Error())
		}
		return ngamux.JSON(rw, in)
	})

	http.ListenAndServe(":8080", mux)
}
