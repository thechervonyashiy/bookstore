package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thechervonyashiy/bookstore/internal/handlers"
	mid "github.com/thechervonyashiy/bookstore/internal/middleware"
)

func SetupRoutes(handler *handlers.Handler) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(mid.JSONMiddleware)

	// Books routes
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetAllBooks) // GET /books
		r.Post("/", handler.CreateBook) // POST /books

		r.Route("/{bookID}", func(r chi.Router) {
			r.Get("/", handler.GetBookByID)   // GET /books/{bookID}
			r.Put("/", handler.UpdateBook)    // PUT /books/{bookID}
			r.Delete("/", handler.DeleteBook) // DELETE /books/{bookID}
		})
	})

	return r
}
