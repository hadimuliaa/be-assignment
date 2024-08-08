package models

// Transaction godoc
type Transaction struct {
    BaseModel
    Amount    float64
    Timestamp int64
    ToAddress string
    Status    string
    AccountID uint      `gorm:"index"` 
    Account   Account   `gorm:"-:all"`
}
