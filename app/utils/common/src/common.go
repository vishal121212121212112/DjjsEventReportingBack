package common

import (
	constants "event-reporting/app/utils/constants"
	"event-reporting/app/utils/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateJSONResponse(c *gin.Context, status constants.ResponseType, statusCode int, message string, data interface{}) {
	var jsonResponse gin.H

	if status == constants.ResponseOK {
		jsonResponse = gin.H{
			"status": status,
			"data":   data,
		}
	} else {
		jsonResponse = gin.H{
			"status":  status,
			"message": message,
		}
	}
	c.JSON(statusCode, jsonResponse)
}

func DecryptAndExtractUserInfo(token string) (string, string, error) {
	// Initialize JWT Service
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		return "", "", fmt.Errorf("JWT service initialization failed")
	}

	// Decrypt the token
	decryptClaims, err := jwtService.Decrypt(token, "")
	if err != nil {
		return "", "", fmt.Errorf("invalid token")
	}

	// Extract user data from token
	dataMap, ok := decryptClaims["data"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid token format: missing data")
	}

	// Get userType (user/employee)
	userType, ok := dataMap["userType"].(string)
	if !ok || userType == "" {
		return "", "", fmt.Errorf("invalid token format: missing data")
	}

	// Get userID or employeeID based on userType
	var uniqueID string
	if userType == "COMPANY" {
		uniqueID, _ = dataMap["userID"].(string)
	} else if userType == "EMPLOYEE" {
		uniqueID, _ = dataMap["employeeID"].(string)
	}

	if uniqueID == "" {
		return "", "", fmt.Errorf(
			"invalid or missing user ID for token of type %q",
			userType,
		)
	}

	return userType, uniqueID, nil
}
