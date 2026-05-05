package main

import (
	"log"
	"net/http"

	"gosecuretransfer/internal/handlers"
)

// ✅ CORS middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

func main() {

	http.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /")
		w.Write([]byte("Backend running 🚀"))
	}))

	http.HandleFunc("/login", enableCORS(handlers.LoginHandler))
	http.HandleFunc("/upload", enableCORS(handlers.UploadHandler))
	http.HandleFunc("/download", enableCORS(handlers.DownloadHandler))
	http.HandleFunc("/files", enableCORS(handlers.ListFilesHandler))

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}