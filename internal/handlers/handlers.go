package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thechervonyashiy/bookstore/internal/services"
)

type Handler struct {
	Service services.BookService
}

type BookPayload struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetAllBooks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve books: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode books: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	bookIDStr := chi.URLParam(r, "bookID")

	id, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.Service.GetBookByID(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload BookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if payload.Title == "" || payload.Author == "" {
		http.Error(w, "Title and Author are required", http.StatusBadRequest)
		return
	}

	id, err := h.Service.CreateBook(payload.Title, payload.Author)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Book created successfully",
	})
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Обновить книгу"))
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	bookIDStr := chi.URLParam(r, "bookID")

	// Преобразуем ID из строки в int
	id, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	cnt, err := h.Service.DeleteBook(int64(id))
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": cnt,
	})
}
