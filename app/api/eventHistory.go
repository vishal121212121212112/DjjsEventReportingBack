package api

import (
	eventHistoryHandler "event-reporting/app/handler/eventHistory"
	middleware "event-reporting/app/helpers/middleware"
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
	r.RouterGroup.POST("/eventHistory/post", middleware.JWT(), eventHistoryHandler.EventHistoryPostHandler)
}
