package handlers

import "net/http"

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Получить все книги"))
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}
