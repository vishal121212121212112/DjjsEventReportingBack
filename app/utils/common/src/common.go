package common
import (
	
"github.com/gin-gonic/gin"
	constants "event-reporting/app/utils/constants"

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