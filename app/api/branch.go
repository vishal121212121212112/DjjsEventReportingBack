package api

import (
	branchHandler "event-reporting/app/handler/branch"
	middleware "event-reporting/app/helpers/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

type branchGroup struct {
	RouterGroup *gin.RouterGroup
}

func (r *branchGroup) Init() {
	defer func() {
		fmt.Println("branch API has been initialized")
	}()

	// Public route for branch search
	r.RouterGroup.GET("/branches/search", middleware.JWT(), branchHandler.SearchBranches)

	// Protected routes (if needed)
	// r.RouterGroup.POST("/branches", middleware.JWT(), branchHandler.CreateBranch)
	// r.RouterGroup.PUT("/branches/:id", middleware.JWT(), branchHandler.UpdateBranch)
}
