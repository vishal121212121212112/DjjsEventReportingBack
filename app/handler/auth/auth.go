package auth

// import (
// 	"event-reporting/app/helpers/logger"
// 	"event-reporting/app/helpers/response"
// 	"event-reporting/app/models"
// 	common "event-reporting/app/utils/common/src"
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// func AuthHandler(c *gin.Context) {
// 	logger.Log.Info.Println("app:handlers:authentication:TokenAuthHandler")

// 	// Parse request body
// 	var newToken models.Auth
// 	if err := c.ShouldBindJSON(&newToken); err != nil {
// 		log.Printf("Error binding JSON: %v", err)
// 		response.SendBadRequestResponse(c, "Invalid request data: "+err.Error())
// 		return
// 	}

// 	token := newToken.Token

// 	userType, uniqueID, _ := common.DecryptAndExtractUserInfo(token)

// 	// Validate User
// 	if userType == "COMPANY" {
// 		res, err := userService.ValidateUser(uniqueID)
// 		if err != nil {
// 			log.Println("User validation failed:", err)
// 			response.SendBadRequestResponse(c, "User validation failed")
// 			return
// 		}
// 		response.SendSuccessResponse(c, res)
// 		return
// 	}

// 	// Validate Employee
// 	if userType == "EMPLOYEE" {
// 		res, err := employeeService.ValidateEmployee(uniqueID)
// 		if err != nil {
// 			log.Println("Employee validation failed:", err)
// 			response.SendBadRequestResponse(c, "Employee validation failed")
// 			return
// 		}
// 		response.SendSuccessResponse(c, res)
// 		return
// 	}
// 	response.SendUnauthorizedResponse(c, "Invalid Token")
// }
