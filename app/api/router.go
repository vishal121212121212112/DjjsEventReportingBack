package api

import (
	"net/http"

	"event-reporting/app/handlers/user"
	"event-reporting/app/helpers/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	User      *handlers.UserHandler
	JWTSecret string
}

func NewRouter(rt Routes) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.RequestID(), gin.Logger())

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

	api := r.Group("/api")
	rt.User.RegisterRoutes(api.Group("/users"))

	// example protected group
	// auth := api.Group("/", middleware.JWT([]byte(rt.JWTSecret)))
	// auth.GET("/me", ...)

	return r
}
