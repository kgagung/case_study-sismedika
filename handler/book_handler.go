package handler

import (
	"Sismedika/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Simpan data buku dalam slice
var books []model.Book = []model.Book{
	{ID: 1, Title: "Book One", Author: "Handoyo", PublishedYear: 2020},
	{ID: 2, Title: "Book Two", Author: "Albert Budi Christian", PublishedYear: 2021},
	{ID: 3, Title: "Book Three", Author: "FX Padmanto Kristiana", PublishedYear: 2022},
	{ID: 4, Title: "Book Four", Author: "Muhammad Febrian Ardiansyah", PublishedYear: 2023},
	{ID: 5, Title: "Book Five", Author: "Heri Ju Abdin Sada", PublishedYear: 2024},
}

// GetBooksHandler mengembalikan daftar semua buku
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// CreateBookHandler menambahkan buku baru ke dalam daftar
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

// GetBookIDHandler mengembalikan buku berdasarkan ID
func GetBookIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, book := range books {
		if book.ID == bookID {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// UpdateBookHandler memperbarui data buku berdasarkan ID
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == bookID {
			books[i] = updatedBook
			books[i].ID = bookID
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// DeleteBookHandler menghapus buku berdasarkan ID
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == bookID {
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}