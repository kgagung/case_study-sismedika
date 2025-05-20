package router

import (
	"Sismedika/handler"
	"Sismedika/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

// NewRouter mengembalikan router baru dengan semua rute yang telah ditentukan
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	// Middleware
	router.Use(
		chiMiddleware.RequestID,
		chiMiddleware.RealIP,
		middleware.LoggerMiddleware,
	)

	// Rute untuk buku
	router.Get("/books", handler.GetBooksHandler)
	router.Post("/books", handler.CreateBookHandler)
	router.Get("/books/{id}", handler.GetBookIDHandler)
	router.Put("/books/{id}", handler.UpdateBookHandler)
	router.Delete("/books/{id}", handler.DeleteBookHandler)

	return router
}