package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gui-laranjeira/go-cleanarch/configs"
	"github.com/gui-laranjeira/go-cleanarch/internal/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	r := chi.NewRouter()
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handlers.GetAllBooksHandler)
		r.Get("/{author}", handlers.GetBookByAuthorHandler)
		r.Post("/", handlers.CreateBookHandler)
	})

	fmt.Println("Server is running on port :" + configs.GetPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort()), r)
}
