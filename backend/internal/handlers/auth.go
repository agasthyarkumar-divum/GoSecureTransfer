package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gosecuretransfer/internal/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🔐 Login request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// ⚠️ Temporary validation (no DB yet)
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		log.Println("❌ Token generation failed:", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}