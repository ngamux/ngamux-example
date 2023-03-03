package main

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()
	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).Text("GET /")
	})

	mux.Post("/", func(rw http.ResponseWriter, r *http.Request) error {
		c := ngamux.NewCtx(rw, r)

		in := map[string]string{}
		err := c.Req.JSON(&in)
		if err != nil {
			c.Res.Status(http.StatusBadGateway).Json(err.Error())
		}
		return c.Res.Json(in)
	})

	http.ListenAndServe(":8080", mux)
}
