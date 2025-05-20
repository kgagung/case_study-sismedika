package main

import (
	"Sismedika/router"
	"log"
	"net/http"
)

// main adalah titik masuk aplikasi
func main() {
	router := router.NewRouter()
	log.Println("Server berjalan di http://localhost:8080/books")
	http.ListenAndServe(":8080", router)
}