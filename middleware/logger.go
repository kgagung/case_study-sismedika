package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

// LoggerMiddleware adalah middleware untuk mencatat setiap permintaan yang masuk
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Buka file log (append mode)
		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Println("Gagal membuka log.txt:", err)
			return
		}
		defer f.Close()

		// Set log output ke file
		log.SetOutput(f)

		start := time.Now()

		// Jalankan handler berikutnya
		next.ServeHTTP(w, r)

		// Tulis log ke file
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
			r.UserAgent(),
		)
	})
}
