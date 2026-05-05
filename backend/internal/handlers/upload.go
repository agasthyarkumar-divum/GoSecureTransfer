package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/crypto"
	"gosecuretransfer/internal/storage"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📤 Upload request received")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)

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

	// 📁 3. Read file
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("❌ Error reading file:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Println("📁 File:", header.Filename)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("❌ Failed to read file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔐 4. Encrypt file
	encrypted, nonce, err := crypto.Encrypt(data, EncryptionKey)
	if err != nil {
		log.Println("❌ Encryption failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	finalData := append(nonce, encrypted...)

	// 💾 5. Save file
	err = os.WriteFile("storage/"+header.Filename, finalData, 0644)
	if err != nil {
		log.Println("❌ File save failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔐 6. Store ownership
	storage.FileOwner[header.Filename] = user

	log.Println("✅ File encrypted & stored:", header.Filename)

	// 📤 Response
	fmt.Fprintf(w, "Uploaded: %s", header.Filename)
}