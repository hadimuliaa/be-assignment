package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database
var DB *gorm.DB

// Connect menghubungkan ke database PostgreSQL menggunakan GORM
func Connect() {
    log.Println("Memuat variabel lingkungan dari file .env")
    // Memuat variabel lingkungan dari file .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Gagal memuat file .env")
    }

    log.Println("Mengambil DATABASE_URL dari variabel lingkungan")
    // Mendapatkan DATABASE_URL dari variabel lingkungan
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL tidak diset dalam variabel lingkungan")
    }

    log.Println("DATABASE_URL:", dsn)

    log.Println("Menghubungkan ke database")
    // Menghubungkan ke database
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal menghubungkan ke database: ", err)
    }

    log.Println("Koneksi ke database berhasil")
}
