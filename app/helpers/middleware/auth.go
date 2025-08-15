package middleware

import (
	"event-reporting/app/helpers/response"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthHandler(c *gin.Context) {
	log.Println("app:handlers:authentication:AuthHandler")

	var req struct {
		Token string `json:"token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequestResponse(c, "Invalid request data: "+err.Error())
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		response.SendErrorResponse(c, 500, "JWT secret not configured")
		return
	}

	token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		response.SendUnauthorizedResponse(c, "Invalid Token")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		response.SendUnauthorizedResponse(c, "Invalid Token Claims")
		return
	}

	userID, _ := claims["user_id"].(string)
	userType, _ := claims["type"].(string)

	if userID == "" || userType == "" {
		response.SendUnauthorizedResponse(c, "Invalid Token Payload")
		return
	}

	// You can add more validation here if needed (e.g., check user in DB)

	response.SendSuccessResponse(c, gin.H{
		"user_id": userID,
		"type":    userType,
		"message": "Token is valid",
	})
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("JWT_SECRET")
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) { return secret, nil })
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userID, _ := claims["user_id"].(string)
		userType, _ := claims["type"].(string)
		c.Set("user_id", userID)
		c.Set("user_type", userType)
		c.Next()
	}
}
