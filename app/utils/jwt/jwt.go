package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey string
}

type JWTClaims struct {
	Data any `json:"data"`
	jwt.RegisteredClaims
}

func NewJWTService() (*JWTService, error) {
	secretKey := os.Getenv("TOKEN_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("TOKEN_SECRET_KEY is not set in environment variables")
	}
	return &JWTService{secretKey: secretKey}, nil
}

func (service *JWTService) GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		Data: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        userID,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (e *JWTService) Encrypt(data map[string]interface{}, expirationDuration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expirationDuration)
	claims := &JWTClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(e.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service *JWTService) Decrypt(tokenString string, secret string) (map[string]interface{}, error) {
	tokenSecret := service.secretKey
	if secret != "" {
		tokenSecret = secret
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data := make(map[string]interface{})
		for key, value := range claims {
			data[key] = value
		}
		return data, nil
	}
	return nil, errors.New("invalid token")
}
