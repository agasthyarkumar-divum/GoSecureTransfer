package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"gosecuretransfer/internal/auth"
	"gosecuretransfer/internal/db"
)

// DeleteFileHandler deletes a single file
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🗑️ Delete file request received")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)

	// 🔍 Only allow DELETE method
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	log.Println("📁 File to delete:", filename)

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
		log.Println("❌ Unauthorized delete attempt")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// 💾 4. Delete from database
	result, err := db.DB.Exec(
		"DELETE FROM files WHERE filename = $1 AND owner = $2",
		filename,
		user,
	)
	if err != nil {
		log.Println("❌ Database delete failed:", err)
		http.Error(w, "Failed to delete file metadata: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Println("❌ No rows deleted")
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}

	// 📁 5. Delete from filesystem
	err = os.Remove("storage/" + filename)
	if err != nil {
		log.Println("⚠️ File not found on filesystem:", err)
		// Don't error out - database is already updated
	}

	log.Println("✅ File deleted:", filename)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "File deleted successfully",
		"filename": filename,
	})
}

// DeleteAllFilesHandler deletes all files for a user
func DeleteAllFilesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🗑️ Delete all files request received")

	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)

	// 🔍 Only allow DELETE method
	if r.Method != http.MethodDelete {
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

	// 📁 3. Get all files for user
	rows, err := db.DB.Query(
		"SELECT filename FROM files WHERE owner = $1",
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

	// 💾 4. Delete all from database
	result, err := db.DB.Exec(
		"DELETE FROM files WHERE owner = $1",
		user,
	)
	if err != nil {
		log.Println("❌ Database delete failed:", err)
		http.Error(w, "Failed to delete files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("❌ Error getting rows affected:", err)
	}

	// 📁 5. Delete all from filesystem
	deletedCount := 0
	for _, filename := range filenames {
		err = os.Remove("storage/" + filename)
		if err == nil {
			deletedCount++
		} else {
			log.Println("⚠️ File not found on filesystem:", filename)
		}
	}

	log.Println("✅ All files deleted for user:", user, "- Count:", rowsAffected)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "All files deleted successfully",
		"deleted": rowsAffected,
	})
}
