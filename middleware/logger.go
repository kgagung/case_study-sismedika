package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware adalah middleware untuk mencatat setiap permintaan yang masuk
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Menjalankan handler berikutnya
		next.ServeHTTP(w, r)

		// Mencatat informasi permintaan
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
			r.UserAgent(),
		)
	})
}