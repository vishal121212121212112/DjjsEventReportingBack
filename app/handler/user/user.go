package userHandler

import (
	"event-reporting/app/helpers/response"
	"event-reporting/app/models"
	userServiceHandler "event-reporting/app/services/users"
	"event-reporting/app/utils/hashing"
	"event-reporting/app/utils/jwt"
	"fmt"
	"log"

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
	if err := userService.GetUserByEmail(req.Email, &user); err != nil {
		response.SendUnauthorizedResponse(c, "Invalid email or password")
		return
	}

	if !hashing.HashVerify(req.Password, user.Password) {
		response.SendUnauthorizedResponse(c, "Invalid email or password")
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

	// Store token in DB (optional)
	if err := userService.UpdateUserToken(user.ID, token); err != nil {
		log.Println("Failed to update user token:", err)
	}

	c.JSON(200, models.LoginResponse{
		Token:  token,
		UserID: user.ID.String(),
		Type:   user.Type,
	})
}
