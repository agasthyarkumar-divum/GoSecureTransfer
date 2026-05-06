package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/handlers"
	"gosecuretransfer/internal/db"
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
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	// Initialize modules with environment variables
	handlers.InitHandlers()

	auth.InitAuth()

	db.InitDB()

	http.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /")
		w.Write([]byte("Backend running 🚀"))
	}))

	http.HandleFunc("/login", enableCORS(handlers.LoginHandler))
	http.HandleFunc("/upload", enableCORS(handlers.UploadHandler))
	http.HandleFunc("/download", enableCORS(handlers.DownloadHandler))
	http.HandleFunc("/files", enableCORS(handlers.ListFilesHandler))
	http.HandleFunc("/register", enableCORS(handlers.RegisterHandler))
	http.HandleFunc("/delete", enableCORS(handlers.DeleteFileHandler))
	http.HandleFunc("/delete-all", enableCORS(handlers.DeleteAllFilesHandler))

	// Read port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}