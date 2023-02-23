package main

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()
	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).String("GET /")
	})

	mux.Post("/", func(rw http.ResponseWriter, r *http.Request) error {
		res := ngamux.Res(rw)

		in := map[string]string{}
		err := ngamux.GetJSON(r, &in)
		if err != nil {
			res.Status(http.StatusBadGateway).JSON(err.Error())
		}
		return ngamux.Res(rw).JSON(in)
	})

	http.ListenAndServe(":8080", mux)
}
