package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"event-reporting/app/api"
	"event-reporting/cmd/setup"
)

func main() {
	// App setup (configuration, DB, RMQ, etc.)
	configData := setup.AppSetup()

	router := gin.Default()
	router.Static("/files", "./public")


	// Initialize all API routes
	MyRouters := api.Routers{
		Router: router,
	}
	MyRouters.Init()

	// Start HTTP server
	serverAddress := configData.App.Host + ":" + configData.App.Port
	if err := router.Run(serverAddress); err != nil {
		log.Fatal("❌ Unable to start the server:", err)
	}
}
