package handler

import (
	"Sismedika/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetBooksHandler mengembalikan daftar semua buku
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := model.GetBookStore().GetBooks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBookIDHandler mengembalikan buku berdasarkan ID
func GetBookIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	book, found := model.GetBookStore().GetBookByID(id)
	if !found {
		http.Error(w, "Buku tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBookHandler menambahkan buku baru
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Gagal membaca data buku", http.StatusBadRequest)
		return
	}

	model.GetBookStore().AddBook(book)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBookHandler memperbarui buku berdasarkan ID
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var updatedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Gagal membaca data buku", http.StatusBadRequest)
		return
	}

	if !model.GetBookStore().UpdateBook(id, updatedBook) {
		http.Error(w, "Buku tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBookHandler menghapus buku berdasarkan ID
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	if !model.GetBookStore().DeleteBook(id) {
		http.Error(w, "Buku tidak ditemukan", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}