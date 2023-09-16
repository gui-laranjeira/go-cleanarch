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
	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Post("/books", handlers.CreateBookHandler)
	fmt.Println("Server is running on port :" + configs.GetPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort()), r)
}
