package models

type Auth struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"userId"`
	Type   string `json:"type"`
}
