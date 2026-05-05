package handlers

import (
	"log"
	"net/http"
	"os"

	"gosecuretransfer/internal/crypto"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Download request received")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	filename := r.URL.Query().Get("file")
	log.Println("📁 Requested file:", filename)

	data, err := os.ReadFile("storage/" + filename)
	if err != nil {
		log.Println("❌ File not found:", err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Extract nonce (first 12 bytes)
	nonce := data[:12]
	ciphertext := data[12:]

	// 🔐 Decrypt
	decrypted, err := crypto.Decrypt(ciphertext, nonce, key)
	if err != nil {
		log.Println("❌ Decryption failed:", err)
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}

	log.Println("✅ File decrypted:", filename)

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Write(decrypted)
}