package api

import (
	eventHistoryHandler "event-reporting/app/handlers/eventHistory"
	"fmt"

	"github.com/gin-gonic/gin"
)

type eventHistoryGroup struct {
	RouterGroup *gin.RouterGroup
}

func (r *eventHistoryGroup) Init() {
	defer func() {
		fmt.Println("eventHistory API has been initialized")
	}()
	r.RouterGroup.POST("/eventHistory/post", eventHistoryHandler.EventHistoryPostHandler)
}
