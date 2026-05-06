package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/crypto"
	"gosecuretransfer/internal/db"
	"gosecuretransfer/internal/storage"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📤 Upload request received")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)

	// 🔍 Only allow POST method
	if r.Method != http.MethodPost {
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

	// � 6. Insert file metadata into database
	_, err = db.DB.Exec(
		"INSERT INTO files (filename, owner) VALUES ($1, $2)",
		header.Filename,
		user,
	)
	if err != nil {
		log.Println("❌ Database insert failed:", err)
		http.Error(w, "Failed to store file metadata: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔐 7. Store ownership in memory (backup)
	storage.FileOwner[header.Filename] = user

	log.Println("✅ File encrypted & stored:", header.Filename)

	// 📤 Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "File uploaded successfully",
		"filename": header.Filename,
	})
}