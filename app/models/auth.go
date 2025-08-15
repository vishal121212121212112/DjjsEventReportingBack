package models

type Auth struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"` // Can be email or username
	Password   string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"userId"`
	Type   string `json:"type"`
}
