package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/crypto"
	"gosecuretransfer/internal/db"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Download request received")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)

	// 🔍 Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Query().Get("file")
	log.Println("📁 Requested file:", filename)

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

	// 🔐 3. Check ownership from database
	var owner string
	err = db.DB.QueryRow(
		"SELECT owner FROM files WHERE filename = $1",
		filename,
	).Scan(&owner)
	if err != nil {
		log.Println("❌ File not found in database:", err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	if owner != user {
		log.Println("❌ Unauthorized access attempt")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// 📁 4. Read file
	data, err := os.ReadFile("storage/" + filename)
	if err != nil {
		log.Println("❌ File not found:", err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// 🔐 5. Decrypt
	nonce := data[:12]
	ciphertext := data[12:]

	decrypted, err := crypto.Decrypt(ciphertext, nonce, EncryptionKey)
	if err != nil {
		log.Println("❌ Decryption failed:", err)
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}

	log.Println("✅ File served:", filename)

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Write(decrypted)
}