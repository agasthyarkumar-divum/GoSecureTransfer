package main

import (
	"log"
	"net/http"

	"gosecuretransfer/internal/handlers"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Backend running 🚀"))
	})

	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)
	http.HandleFunc("/files", handlers.ListFilesHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}