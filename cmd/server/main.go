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
	r.Get("/api/v1/ping", handlers.HealthCheckHandler)

	r.Route("/api/v1/books", func(r chi.Router) {
		r.Get("/", handlers.GetAllBooksHandler)
		r.Get("/{id}", handlers.GetBookByIDHandler)
		r.Post("/", handlers.CreateBookHandler)
		r.Delete("/{id}", handlers.DeleteBookHandler)
		r.Put("/{id}", handlers.UpdateBookHandler)
	})

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", handlers.GetAllUsersHandler)
		r.Post("/", handlers.CreateUserHandler)
		r.Get("/{id}", handlers.GetUserByIDHandler)
		r.Put("/{id}", handlers.UpdateUserHandler)
		r.Patch("/secret/{id}", handlers.UpdatePasswordHandler)
	})

	fmt.Println("Server is running on port :" + configs.GetPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort()), r)
}
