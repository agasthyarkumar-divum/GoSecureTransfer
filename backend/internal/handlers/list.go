package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/db"
)

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📂 List files request")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)
	w.Header().Set("Content-Type", "application/json")

	// 🔍 Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// 📁 3. Query files from database
	rows, err := db.DB.Query(
		"SELECT filename FROM files WHERE owner = $1 ORDER BY id DESC",
		user,
	)
	if err != nil {
		log.Println("❌ Failed to query files:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var filenames []string
	for rows.Next() {
		var filename string
		if err := rows.Scan(&filename); err != nil {
			log.Println("❌ Error scanning file:", err)
			continue
		}
		filenames = append(filenames, filename)
	}

	if err := rows.Err(); err != nil {
		log.Println("❌ Error iterating files:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("📁 Files for user:", filenames)

	json.NewEncoder(w).Encode(filenames)
}