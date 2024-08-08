package controllers

import (
    "account-manager/config"
    "account-manager/models"
    "net/http"
    "time"
    "log"

    "github.com/gin-gonic/gin"
)
// SendTransaction godoc
// @Summary Kirim transaksi baru
// @Description Tambahkan transaksi baru dan proses secara asinkron
// @Tags transaksi
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Data Transaksi"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /send [post]
func SendTransaction(c *gin.Context) {
    var transaction models.Transaction
    
    // Bind JSON input to the transaction struct
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Log the incoming transaction details for debugging
    log.Printf("Received transaction: %+v", transaction)

    // Add timestamp and initial status
    transaction.Timestamp = time.Now().Unix()
    transaction.Status = "pending"

    // Log the transaction details before saving to the database
    log.Printf("Saving transaction: %+v", transaction)

    // Attempt to save the transaction to the database
    if err := config.DB.Create(&transaction).Error; err != nil {
        log.Printf("Error saving transaction: %v", err) // Log error if saving fails
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Process the transaction asynchronously
    go processTransaction(transaction)

    // Return a successful response
    c.JSON(http.StatusOK, gin.H{"message": "Transaction initiated", "transaction": transaction})
}

func processTransaction(transaction models.Transaction) {
    time.Sleep(30 * time.Second)
    transaction.Status = "completed"

    // Save the updated transaction status
    config.DB.Save(&transaction)
}
// -------------------
// WithdrawTransaction godoc
// @Summary Tarik dana dari akun
// @Description Tambahkan transaksi penarikan baru dan proses secara asinkron
// @Tags transaksi
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Data Transaksi"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /withdraw [post]
func WithdrawTransaction(c *gin.Context) {
    var transaction models.Transaction

    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Printf("Received withdrawal request: %+v", transaction)

    transaction.Timestamp = time.Now().Unix()
    transaction.Status = "pending"

    log.Printf("Saving withdrawal transaction: %+v", transaction)

    if err := config.DB.Create(&transaction).Error; err != nil {
        log.Printf("Error saving withdrawal transaction: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    go processWithdrawTransaction(transaction)

    c.JSON(http.StatusOK, gin.H{"message": "Withdrawal initiated", "transaction": transaction})
}

func processWithdrawTransaction(transaction models.Transaction) {
    time.Sleep(30 * time.Second)
    transaction.Status = "completed"

    // Save the updated transaction status
    config.DB.Save(&transaction)
}

// ----handler utk get
// GetTransactions godoc
// @Summary Ambil semua transaksi
// @Description Ambil semua transaksi dari database
// @Tags transaksi
// @Produce json
// @Success 200 {array} models.Transaction
// @Failure 500 {object} models.ErrorResponse
// @Router /transactions [get]
func GetTransactions(c *gin.Context) {
    var transactions []models.Transaction
    if err := config.DB.Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, transactions)
}

// GetTransactionsByAccountID godoc
// @Summary Ambil semua transaksi berdasarkan ID akun
// @Description Ambil semua transaksi dari database berdasarkan ID akun
// @Tags transaksi
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {array} models.Transaction
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /accounts/{id}/transactions [get]
func GetTransactionsByAccountID(c *gin.Context) {
    var transactions []models.Transaction
    accountID := c.Param("id")

    if err := config.DB.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, transactions)
}