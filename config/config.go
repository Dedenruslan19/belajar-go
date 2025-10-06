package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file. Using environment variables from the OS.")
	}
}

func ConnectDB() *sql.DB {
    LoadEnv() 

    // 2. Ambil Variabel dari Environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

    // 3. Validasi Variabel
	if host == "" || port == "" || user == "" || pass == "" || name == "" {
		log.Fatal("❌ Error: One or more required database environment variables (DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME) are not set.")
	}
    
    // 4. Buat Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

    // 5. Buka Koneksi Database (tidak benar-benar terhubung)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Fatal Error: Failed to open database connection: %v", err)
	}

    // 6. Tes Koneksi (Ping)
	err = db.Ping()
	if err != nil {
		// Jika Ping gagal, matikan program dan tampilkan error
		log.Fatalf("❌ Fatal Error: Failed to connect to database. Check credentials or service status. Error: %v", err)
	}

	fmt.Println("✅ Database successfully connected!")
	return db
}