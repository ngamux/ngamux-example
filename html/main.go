package main

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	m := ngamux.New()

	m.Get("/", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).Html("./html/index.html", ngamux.Map{
			"title":   "index",
			"message": "welcome to our site :)",
		})
	})

	http.ListenAndServe(":8080", m)
}
