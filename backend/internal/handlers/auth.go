package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/storage"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 🔐 REGISTER
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🆕 Register request")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)

	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Invalid input", 400)
		return
	}

	// ❌ Check if user already exists
	if _, exists := storage.Users[req.Username]; exists {
		http.Error(w, "Username already exists", 400)
		return
	}

	// 🔐 Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Hashing error", 500)
		return
	}

	// 💾 Store user
	storage.Users[req.Username] = string(hash)

	log.Println("✅ User registered:", req.Username)

	w.Write([]byte("User registered successfully"))
}

// 🔐 LOGIN
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🔐 Login request")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)
	w.Header().Set("Content-Type", "application/json")

	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	storedHash, exists := storage.Users[req.Username]
	if !exists {
		http.Error(w, "User not found", 401)
		return
	}

	// 🔐 Compare password
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid password", 401)
		return
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "Token error", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}