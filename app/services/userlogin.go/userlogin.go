package userlogin

// import (
// 	"event-reporting/app/helpers/response"
// 	"event-reporting/app/models"

// 	"github.com/gin-gonic/gin"
// )

// func LoginHandler(c *gin.Context) {
// 	var req models.LoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		response.SendBadRequestResponse(c, "Invalid input: "+err.Error())
// 		return
// 	}

// 	user, err := userService.AuthenticateUser(req.Email, req.Password)
// 	if err != nil {
// 		response.SendUnauthorizedResponse(c, "Invalid email or password")
// 		return
// 	}

// 	// Generate JWT token (implement your own logic)
// 	token, err := authService.GenerateToken(user)
// 	if err != nil {
// 		response.SendInternalServerError(c, "Failed to generate token")
// 		return
// 	}

// 	c.JSON(200, models.LoginResponse{
// 		Token:  token,
// 		UserID: user.ID.String(),
// 		Type:   user.Type,
// 	})
// }
