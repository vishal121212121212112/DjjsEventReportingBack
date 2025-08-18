package main

import (
	"event-reporting/app/api"
	"event-reporting/cmd/setup"
	"log"

	"github.com/gin-gonic/gin"
)

// func SeedAdmin(repo *database.Repository) error {
// 	adminEmail := "admin@djjs.org"
// 	var user models.User
// 	err := repo.Find(&user, map[string]interface{}{"email": adminEmail})
// 	if err == nil {
// 		return nil // already exists
// 	}
// 	// Hash the password
// 	hashedPassword, err := hashing.HashData("admin123")
// 	if err != nil {
// 		return err
// 	}

// 	admin := models.User{
// 		ID:        uuid.New(),
// 		Username:  "admin",
// 		Email:     adminEmail,
// 		Password:  hashedPassword, // Use hashed password
// 		Type:      "hoadmin",
// 		CreatedOn: time.Now().Format(time.RFC3339),
// 		UpdatedOn: time.Now().Format(time.RFC3339),
// 	}
// 	return repo.Create(&admin)
// }

func main() {
	// App setup (configuration, DB, RMQ, etc.)
	configData := setup.AppSetup()

	router := gin.Default()
	router.Static("/files", "./public")

	// Initialize all API routes
	MyRouters := api.Routers{
		Router: router,
	}
	// var repo *database.Repository
	// SeedAdmin(repo)
	MyRouters.Init()

	// Start HTTP server
	serverAddress := configData.App.Host + ":" + configData.App.Port
	if err := router.Run(serverAddress); err != nil {
		log.Fatal("❌ Unable to start the server:", err)
	}
}
