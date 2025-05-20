package main

import (
	"Sismedika/model"
	"Sismedika/router"
	"log"
	"net/http"
)

// data dari bookstore
func init() {
	// Inisialisasi data awal untuk BookStore
	store := model.GetBookStore()
	store.AddBook(model.Book{Title: "Book One", Author: "Handoyo", PublishedYear: 2020})
	store.AddBook(model.Book{Title: "Book Two", Author: "Albert Budi Christian", PublishedYear: 2019})
	store.AddBook(model.Book{Title: "Book Three", Author: "FX Padmanto Kristiana", PublishedYear: 2021})
	store.AddBook(model.Book{Title: "Book Four", Author: "Muhammad Febrian Ardiansyah", PublishedYear: 2018})
	store.AddBook(model.Book{Title: "Book Five", Author: "Heri Ju Abdin Sada", PublishedYear: 2022})
}
// main adalah titik masuk aplikasi
func main() {
	router := router.NewRouter()
	log.Println("Server berjalan di http://localhost:8080/books")
	http.ListenAndServe(":8080", router)
}