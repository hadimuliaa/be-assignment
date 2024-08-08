package controllers

import (
    "account-manager/config"
    "account-manager/models"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// CreateAutoPayment godoc
// @Summary Buat pembayaran otomatis
// @Description Buat pembayaran otomatis untuk pengguna
// @Tags pembayaran otomatis
// @Accept json
// @Produce json
// @Param autoPayment body models.AutoPayment true "Data Pembayaran Otomatis"
// @Success 200 {object} models.AutoPayment
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /autopayment [post]
func CreateAutoPayment(c *gin.Context) {
    var autoPayment models.AutoPayment
    if err := c.ShouldBindJSON(&autoPayment); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    autoPayment.NextRun = calculateNextRun(autoPayment.Interval)
    if err := config.DB.Create(&autoPayment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, autoPayment)
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
