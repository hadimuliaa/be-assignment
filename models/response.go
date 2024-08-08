package models

// ResponseMessage godoc
type ResponseMessage struct {
    Message string `json:"message"`
}

// ErrorResponse godoc
type ErrorResponse struct {
    Error string `json:"error"`
}

// TokenResponse godoc
type TokenResponse struct {
    Token string `json:"token"`
}

// LoginRequest godoc
type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
