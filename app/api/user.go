package api

import (
	userHandler "event-reporting/app/handler/user"
	middleware "event-reporting/app/helpers/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

type userGroup struct {
	RouterGroup *gin.RouterGroup
}

func (r *userGroup) Init() {
	defer func() {
		fmt.Println("users api has been intillized")
	}()

	r.RouterGroup.POST("/users", middleware.JWT(), userHandler.CreateUserHandler)
	r.RouterGroup.POST("/users/login", userHandler.LoginHandler)
}
