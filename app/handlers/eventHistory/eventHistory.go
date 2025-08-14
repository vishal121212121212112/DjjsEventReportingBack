package eventHistoryHandler

import (
	"context"
	"event-reporting/app/helpers/logger"
	"event-reporting/app/helpers/response"
	"event-reporting/app/models"
	eventHistoryService "event-reporting/app/services/eventHistory"

	"github.com/gin-gonic/gin"
)

var eventHistoryHandler eventHistoryService.EventHistoryService

//EventHistoryPostHandler
func EventHistoryPostHandler(c *gin.Context) {
	logger.Log.Info.Println("app:handlers:eventHisotry:eventHisotryHandler")
	var newEventHistoryReq models.EventHistory
	if err := c.ShouldBindJSON(&newEventHistoryReq); err != nil {
		logger.Log.Error.Errorln("Error binding JSON:", err)
		response.SendBadRequestResponse(c, "Invalid request data: "+err.Error())
		return
	}
	createEventHistory, err := eventHistoryHandler.CreateEventHistory(context.Background(), &newEventHistoryReq)
	if err != nil {
		response.SendBadRequestResponse(c, err.Error())
		return
	}
	response.SendSuccessResponse(c, createEventHistory)
}
