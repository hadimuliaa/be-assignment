package controllers

import (
    "account-manager/config"
    "account-manager/models"
    "net/http"

    "github.com/gin-gonic/gin"
)
// CreateAccount godoc
// @Summary Membuat akun baru untuk pengguna
// @Description Membuat akun baru untuk pengguna
// @Tags akun
// @Accept json
// @Produce json
// @Param account body models.Account true "Data Akun"
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /accounts [post]
func CreateAccount(c *gin.Context) {
    var account models.Account
    if err := c.ShouldBindJSON(&account); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&account).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})
}
// GetAccounts godoc
// @Summary Ambil semua akun
// @Description Ambil semua akun dari database
// @Tags akun
// @Produce json
// @Success 200 {array} models.Account
// @Failure 500 {object} models.ErrorResponse
// @Router /accounts [get]
func GetAccounts(c *gin.Context) {
    var accounts []models.Account
    if err := config.DB.Find(&accounts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, accounts)
}