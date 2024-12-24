package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Route("/books", func(r chi.Router) {
		r.Get("/", handlers.GetAllBooks)
		r.Post("/", handlers.CreateBook)

		r.Route("/{bookID}", func(r chi.Router) {
			r.Get("/", handlers.GetBookByID)
			r.Put("/", handlers.UpdateBook)
			r.Delete("/", handlers.DeleteBook)
		})
	})

	http.ListenAndServe(":8080", router)
}
