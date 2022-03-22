package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"internshipApplicationTemplate/internal/config"
	"internshipApplicationTemplate/internal/handlers"
	"log"
	"net/http"
)

func main() {
	config.New()

	serverAddress := config.Cfg.ServerAddress
	log.Println("the server address is", serverAddress)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/chartas/", handlers.CreateImage)

	log.Fatal(http.ListenAndServe(serverAddress, router))
}
