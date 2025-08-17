package userHandler

import (
	"event-reporting/app/helpers/response"
	"event-reporting/app/models"
	userServiceHandler "event-reporting/app/services/users"
	"event-reporting/app/utils/hashing"
	"event-reporting/app/utils/jwt"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var userService userServiceHandler.UserService

func CreateUserHandler(c *gin.Context) {
	log.Println("CreateUserHandler: received request to create user")
	var req models.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequestResponse(c, fmt.Sprintf("Invalid input: %v", err.Error()))
		return
	}

	creatorID, exists := c.Get("user_id")
	if !exists {
		response.SendUnauthorizedResponse(c, "Unauthorized: invalid token")
		return
	}

	var creator models.User
	if err := userService.GetUserByID(creatorID.(string), &creator); err != nil {
		response.SendBadRequestResponse(c, "Creator user not found")
		return
	}

	userId, err := userService.CreateUser(creator, req)
	if err != nil {
		response.SendBadRequestResponse(c, fmt.Sprintf("Failed to create user: %v", err.Error()))
		return
	}
	response.SendSuccessResponse(c, gin.H{
		"userId": userId,
	})
}

func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequestResponse(c, "Invalid input: "+err.Error())
		return
	}

	var user models.User
	// Use the new method that supports both email and username
	if err := userService.GetUserByIdentifier(req.Identifier, &user); err != nil {
		response.SendUnauthorizedResponse(c, "Invalid credentials")
		return
	}

	if !hashing.HashVerify(req.Password, user.Password) {
		response.SendUnauthorizedResponse(c, "Invalid credentials")
		return
	}

	jwtService, err := jwt.NewJWTService()
	if err != nil {
		response.SendErrorResponse(c, 500, "JWT service error")
		return
	}

	token, err := jwtService.GenerateToken(user.ID.String(), user.Type, user.Email)
	if err != nil {
		response.SendErrorResponse(c, 500, "Failed to generate token")
		return
	}

	// Update token, login timestamps, and expiration
	currentTime := time.Now().Format(time.RFC3339)
	expirationTime := time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	additionalUpdates := map[string]interface{}{
		"last_login_on": currentTime,
		"updated_on":    currentTime,
		"expired_on":    expirationTime,
	}

	// Set first_login_on only if it's the first login
	if user.FirstLoginOn == "" {
		additionalUpdates["first_login_on"] = currentTime
	}

	if err := userService.UpdateUserToken(user.ID, token, additionalUpdates); err != nil {
		log.Println("Failed to update user login info:", err)
		// Don't fail the login if this update fails
	}

	c.JSON(200, models.LoginResponse{
		Token:  token,
		UserID: user.ID.String(),
		Type:   user.Type,
	})
}
