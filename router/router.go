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
	router.Route("/books", func(router chi.Router) {
		router.Get("/", handler.GetBooksHandler)
		router.Get("/{id}", handler.GetBookIDHandler)
		router.Post("/", handler.CreateBookHandler)
		router.Put("/{id}", handler.UpdateBookHandler)
		router.Delete("/{id}", handler.DeleteBookHandler)
	})

	return router
}