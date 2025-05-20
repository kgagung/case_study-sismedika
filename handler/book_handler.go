package handler

import (
	"Sismedika/model"
	"Sismedika/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetBooksHandler mengembalikan daftar semua buku
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := model.GetBookStore().GetBooks()
	utils.JSONResponse(w, http.StatusOK, books)
}

// GetBookIDHandler mengembalikan buku berdasarkan ID
func GetBookIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	book, found := model.GetBookStore().GetBookByID(id)
	if !found {
		utils.JSONError(w, http.StatusNotFound, "Buku tidak ditemukan")
		return
	}

	utils.JSONResponse(w, http.StatusOK, book)
}

// CreateBookHandler menambahkan buku baru
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Gagal membaca data buku")
		return
	}

	model.GetBookStore().AddBook(book)

	utils.JSONResponse(w, http.StatusCreated, book)
}

// UpdateBookHandler memperbarui buku berdasarkan ID
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var updatedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Gagal membaca data buku")
		return
	}

	if !model.GetBookStore().UpdateBook(id, updatedBook) {
		utils.JSONError(w, http.StatusNotFound, "Buku tidak ditemukan")
		return
	}

	utils.JSONResponse(w, http.StatusOK, updatedBook)
}

// DeleteBookHandler menghapus buku berdasarkan ID
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if !model.GetBookStore().DeleteBook(id) {
		utils.JSONError(w, http.StatusNotFound, "Buku tidak ditemukan")
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, map[string]string{"message": "Buku berhasil dihapus"})
}