package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/gui-laranjeira/go-cleanarch/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/api/v1/ping", handlers.HealthCheckHandler)

	req, err := http.NewRequest("GET", "/api/v1/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "HealthCheckHandler response code should be 200")
}
