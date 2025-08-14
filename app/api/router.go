package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// middleware "event-reporting/app/helpers/middleware"
	adminHandler "event-reporting/app/handlers/admin" 
)

type Routers struct {
	Router *gin.Engine
}

func (r *Routers) Init() {
	// r.Router.Use(middleware.CORSMiddleware())

	// Add health check route
	r.Router.GET("/v1/admin/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Admin Service is up and running",
		})
	})

	r.Router.StaticFile("/admin/swagger.yaml", "./docs/swagger.yaml")

	r.Router.GET("/admin/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/admin/swagger.yaml"),
	))



	v1 := r.Router.Group("/admin")

	// Initialize Admin Routes
	adminGroup := adminHandler.AdminGroup{
		RouterGroup: v1,
	}
	adminGroup.Init()

	defer func() {
		fmt.Println("Router has been initialized..")
	}()
}
