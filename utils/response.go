package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse mengirimkan respons JSON dengan status dan data yang diberikan
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ErrorResponse mengirimkan respons JSON dengan pesan kesalahan
func JSONError(w http.ResponseWriter, status int, message string) {
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}