package main

import (
	"fmt"
	"go-rest-api/configs"
	"go-rest-api/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Starting server...")

	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
