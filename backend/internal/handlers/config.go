package handlers

import (
	"log"
	"os"
)

// EncryptionKey is initialized from environment or default
var EncryptionKey []byte

// CORSOrigin is initialized from environment or default
var CORSOrigin string

// InitHandlers initializes the handlers package with environment variables
func InitHandlers() {
	// Read encryption key from environment
	encKeyStr := os.Getenv("ENCRYPTION_KEY")
	if encKeyStr == "" {
		encKeyStr = "12345678901234567890123456789012" // Default 32-byte key
		log.Println("⚠️ ENCRYPTION_KEY not set, using default")
	}

	if len(encKeyStr) != 32 {
		log.Fatal("❌ ENCRYPTION_KEY must be exactly 32 bytes")
	}

	EncryptionKey = []byte(encKeyStr)

	// Read CORS origin from environment
	CORSOrigin = os.Getenv("CORS_ORIGIN")
	if CORSOrigin == "" {
		CORSOrigin = "*" // Default to allow all origins
	}
}
