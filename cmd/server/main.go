package main

import (
	"net/http"

	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/internal/infra/webservice/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/zipcode/{zipcode}", handlers.GetZipCode)

	http.ListenAndServe(":8000", r)
}
