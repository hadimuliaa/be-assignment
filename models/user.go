package models
import (
    "time"
)

// User godoc
type User struct {
    BaseModel
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Accounts []Account
}
// Account godoc
type Account struct {
    BaseModel
    Type     string
    UserID   uint
    Payments []PaymentHistory
}

// PaymentHistory godoc
type PaymentHistory struct {
    BaseModel
    Amount    float64
    Timestamp int64
    Status    string
    AccountID uint
}
// AutoPayment godoc
type AutoPayment struct {
    BaseModel
    UserID     uint      `json:"user_id"`
    AccountID  uint      `json:"account_id"`
    Amount     float64   `json:"amount"`
    Interval   string    `json:"interval"` // e.g., "daily", "weekly", "monthly"
    NextRun    time.Time `json:"next_run"`
}
