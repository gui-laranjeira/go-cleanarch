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
	r.Route("/api/v1/books", func(r chi.Router) {
		r.Get("/", handlers.GetAllBooksHandler)
		r.Get("/{id}", handlers.GetBookByIDHandler)
		r.Post("/", handlers.CreateBookHandler)
		r.Delete("/{id}", handlers.DeleteBookHandler)
	})

	fmt.Println("Server is running on port :" + configs.GetPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort()), r)
}
