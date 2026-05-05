package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/storage"
)

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📂 List files request")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)
	w.Header().Set("Content-Type", "application/json")

	// 🔐 1. Extract token
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	// 🔐 2. Validate token
	user, err := auth.ValidateToken(tokenStr)
	if err != nil {
		log.Println("❌ Invalid token:", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	log.Println("👤 User:", user)

	// 📁 3. Read storage
	files, err := os.ReadDir("storage")
	if err != nil {
		log.Println("❌ Failed to read storage:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filenames []string

	for _, file := range files {
		name := file.Name()

		// skip hidden files (.keep)
		if name[0] == '.' {
			continue
		}

		// 🔐 Only include files owned by user
		if storage.FileOwner[name] == user {
			filenames = append(filenames, name)
		}
	}

	log.Println("📁 Files for user:", filenames)

	json.NewEncoder(w).Encode(filenames)
}