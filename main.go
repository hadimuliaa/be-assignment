package main

import (
    "account-manager/config"
    "account-manager/models"
    "account-manager/routers"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    _ "account-manager/docs" // Penting untuk import dokumen swagger
    "github.com/robfig/cron/v3"
)

func main() {
    // Muat file .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Koneksi ke database
    config.Connect()

    // Auto-migrate models
    config.DB.AutoMigrate(&models.User{}, &models.Account{}, &models.PaymentHistory{}, &models.Transaction{}, &models.AutoPayment{})

    // Setup cron job untuk pembayaran otomatis
    c := cron.New()
    c.AddFunc("@every 1m", runAutoPayments) // Periksa setiap menit
    c.Start()
    defer c.Stop()

    // Inisialisasi router
    r := routers.SetupRouter()

    // Mendapatkan port dari environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Jalankan server pada port yang ditentukan dan mendengarkan pada semua interface
    log.Printf("Running server on 0.0.0.0:%s...", port)
    if err := r.Run("0.0.0.0:" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

func runAutoPayments() {
    var autoPayments []models.AutoPayment
    now := time.Now()

    if err := config.DB.Where("next_run <= ?", now).Find(&autoPayments).Error; err != nil {
        log.Println("Error fetching auto payments:", err)
        return
    }

    for _, autoPayment := range autoPayments {
        // Lakukan transaksi otomatis
        transaction := models.Transaction{
            AccountID: autoPayment.AccountID,
            Amount:    autoPayment.Amount,
            Timestamp: now.Unix(),
            Status:    "pending",
        }

        if err := config.DB.Create(&transaction).Error; err != nil {
            log.Println("Error creating transaction:", err)
            continue
        }

        // Perbarui waktu berikutnya untuk pembayaran otomatis
        autoPayment.NextRun = calculateNextRun(autoPayment.Interval)
        if err := config.DB.Save(&autoPayment).Error; err != nil {
            log.Println("Error updating auto payment:", err)
        }
    }
}

func calculateNextRun(interval string) time.Time {
    now := time.Now()
    switch interval {
    case "daily":
        return now.Add(24 * time.Hour)
    case "weekly":
        return now.Add(7 * 24 * time.Hour)
    case "monthly":
        return now.AddDate(0, 1, 0)
    default:
        return now
    }
}