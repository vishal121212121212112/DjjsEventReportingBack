package admin

import (
	"github.com/gin-gonic/gin"
	adminHandlers "event-reporting/app/handlers/admin" // Import admin handlers
)

type AdminGroup struct {
	RouterGroup *gin.RouterGroup
}

func (r *AdminGroup) Init() {
	// Register routes for admin operations
	r.RouterGroup.POST("/register", adminHandlers.RegisterAdminHandler)
	r.RouterGroup.POST("/login", adminHandlers.AdminLoginHandler)
	r.RouterGroup.PATCH("/update-password", adminHandlers.UpdateAdminPasswordHandler)
}
