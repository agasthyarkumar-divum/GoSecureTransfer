package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/db"
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

	// 🔍 Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// 💾 Store user in database
	_, err = db.DB.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		req.Username,
		string(hash),
	)
	if err != nil {
		log.Println("❌ Database insert failed:", err)
		http.Error(w, "Failed to register user: "+err.Error(), 500)
		return
	}

	// 💾 Store user in memory (backup)
	storage.Users[req.Username] = string(hash)

	log.Println("✅ User registered:", req.Username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// 🔐 LOGIN
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🔐 Login request")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)
	w.Header().Set("Content-Type", "application/json")
	// 🔍 Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	var storedHash string
	err := db.DB.QueryRow(
		"SELECT password FROM users WHERE username = $1",
		req.Username,
	).Scan(&storedHash)
	if err != nil {
		log.Println("❌ User not found:", err)
		http.Error(w, "User not found", 401)
		return
	}

	// 🔐 Compare password
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password))
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