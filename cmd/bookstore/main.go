package main

import (
	"log"
	"net/http"

	"github.com/thechervonyashiy/bookstore/internal/handlers"
	"github.com/thechervonyashiy/bookstore/internal/routes"
	"github.com/thechervonyashiy/bookstore/internal/services"
	"github.com/thechervonyashiy/bookstore/storage/sqlite"
)

func main() {
	// init db
	storagePath := "./storage.db"
	repo, err := sqlite.New(storagePath)
	if err != nil {
		log.Fatal("failed to init storage ", err)
	}

	bookService := services.NewBookService(repo)

	handler := &handlers.Handler{
		Service: bookService,
	}

	router := routes.SetupRoutes(handler)

	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
