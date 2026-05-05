package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gosecuretransfer/internal/crypto"
)

// ⚠️ Temporary key (32 bytes for AES-256)
var key = []byte("12345678901234567890123456789012")

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📤 Upload request received")

	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	// 🔐 Encrypt
	encrypted, nonce, err := crypto.Encrypt(data, key)
	if err != nil {
		log.Println("❌ Encryption failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	finalData := append(nonce, encrypted...)

	// 💾 Save file
	err = os.WriteFile("storage/"+header.Filename, finalData, 0644)
	if err != nil {
		log.Println("❌ File save failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("✅ File encrypted & stored:", header.Filename)

	fmt.Fprintf(w, "Uploaded: %s", header.Filename)
}