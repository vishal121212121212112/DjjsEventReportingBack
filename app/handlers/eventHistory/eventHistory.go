package eventHistoryHandler

import (
	"context"
	"event-reporting/app/helpers/logger"
	"event-reporting/app/helpers/response"
	"event-reporting/app/models"

	EventHistoryService "event-reporting/app/services/eventHistory"

	"github.com/gin-gonic/gin"
)

var eventHistoryHandler EventHistoryService.EventHistoryService

// EventHistoryPostHandler handles the creation of an event with all related records
func EventHistoryPostHandler(c *gin.Context) {
	logger.Log.Info.Println("app:handlers:eventHistory:EventHistoryPostHandler")

	var req models.CreateFullEventRequest 
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	logger.Log.Error.Errorln("Error binding JSON:", err)
	// 	response.SendBadRequestResponse(c, "Invalid request data: "+err.Error())
	// 	return
	// }

	createdData, err := eventHistoryHandler.CreateFullEvent(context.Background(), &req)
	if err != nil {
		logger.Log.Error.Errorln("Error creating full event:", err)
		response.SendBadRequestResponse(c, err.Error())
		return
	}

	response.SendSuccessResponse(c, createdData)
}