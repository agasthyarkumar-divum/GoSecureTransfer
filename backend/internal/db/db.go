package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection and ensures schema exists
func InitDB() {
	// You can override this via environment variable later
	connStr := os.Getenv("DB_CONN")
	if connStr == "" {
		connStr = "user=postgres password=postgres dbname=securevault sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ DB open error:", err)
	}

	// Verify connection
	if err = DB.Ping(); err != nil {
		log.Fatal("❌ DB ping error:", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	createTables()
}

// createTables creates required tables if they do not exist
func createTables() {
	// USERS TABLE
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		password TEXT NOT NULL
	);
	`)
	if err != nil {
		log.Fatal("❌ Failed to create users table:", err)
	}

	// FILES TABLE
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS files (
		id SERIAL PRIMARY KEY,
		filename TEXT NOT NULL,
		owner TEXT NOT NULL REFERENCES users(username) ON DELETE CASCADE
	);
	`)
	if err != nil {
		log.Fatal("❌ Failed to create files table:", err)
	}

	// INDEX (performance improvement for queries)
	_, err = DB.Exec(`
	CREATE INDEX IF NOT EXISTS idx_files_owner ON files(owner);
	`)
	if err != nil {
		log.Println("⚠️ Failed to create index (non-critical):", err)
	}

	log.Println("✅ Tables ready")
}