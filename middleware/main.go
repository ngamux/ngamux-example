package main

import (
	"fmt"
	"net/http"

	"github.com/ngamux/middleware/ping"
	"github.com/ngamux/ngamux"
)

func globalMiddleware(next ngamux.Handler) ngamux.Handler {
	return func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Println("hello from global middleware")
		return next(rw, r)
	}
}

func routeMiddleware(next ngamux.Handler) ngamux.Handler {
	return func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Println("hello from route middleware")
		return next(rw, r)
	}
}

func main() {
	mux := ngamux.New()
	mux.Use(ping.New())
	mux.Use(globalMiddleware)

	mux.Get("/", routeMiddleware(func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Println("hello / handler")
		return ngamux.Res(rw).Text("GET /")
	}))

	mux.Get("/users", ngamux.WithMiddlewares(routeMiddleware)(func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Println("hello from /users handler")
		return ngamux.Res(rw).Text("GET /users")
	}))

	mux.With(routeMiddleware).Get("/products", func(rw http.ResponseWriter, r *http.Request) error {
		return ngamux.Res(rw).Text("GET /products")
	})

	http.ListenAndServe(":8080", mux)
}
