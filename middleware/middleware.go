package main

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ngamux"
)

func globalMiddleware(next ngamux.HandlerFunc) ngamux.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Println("hello from global middleware")
		return next(rw, r)
	}
}

func routeMiddleware(next ngamux.HandlerFunc) ngamux.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Println("hello from route middleware")
		return next(rw, r)
	}
}

func main() {
	mux := ngamux.NewNgamux(ngamux.Config{
		RemoveTrailingSlash: true,
	})
	mux.Use(globalMiddleware)

	mux.Get("/", routeMiddleware(func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(rw, "hello from / handler")
		fmt.Println("hello / handler")
		return nil
	}))

	mux.Get("/users", ngamux.WithMiddlewares(routeMiddleware)(func(rw http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(rw, "hello from /users handler")
		fmt.Println("hello from /users handler")
		return nil
	}))

	http.ListenAndServe(":8080", mux)
}
