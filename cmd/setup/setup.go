package setup

import (
	"log"
	"sync"

	"event-reporting/app/config"

	"event-reporting/app/database/pgsql/connection"
	"event-reporting/app/helpers/logger"

	"github.com/joho/godotenv"
)

// Initialize all setup processes
func AppSetup() *config.Config {
	setUpLogger()

	configData, err := config.LoadConfigs("config/config.yaml")
	if err != nil {
		log.Fatal("failed to load config file", err)
		panic(err)
	}

	loadEnv(configData.App.Environment)

	

	// Create a WaitGroup to synchronize DB initializations and consumer start
	var wg sync.WaitGroup
	wg.Add(1) // We have two database initializations (MongoDB and PostgreSQL)


	// Initialize PostgreSQL in a goroutine
	go initPostgres(&wg)


	// Wait for both databases to initialize
	wg.Wait()



	return configData
}



func initPostgres(wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	// PostgreSQL connection logic
	connection.ConnectDatabase()
}

func loadEnv(environment string) {
	file := ".env.dev"
	if environment == "production" {
		file = ".env.prod"
	}
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalln("failed to load env file")
		panic(err)
	}

	log.Println("app is running on", environment, "environment")
}

func setUpLogger() {
	loggerOptions := &logger.LoggerSetUpOptions{
		Info: &logger.LoggerOptions{
			Filename:   "./logs/info.log",
			MaxSize:    5,
			MaxBackups: 2,
			MaxAge:     5,
			Compress:   true,
		},
		Error: &logger.LoggerOptions{
			Filename:   "./logs/error.log",
			MaxSize:    5,
			MaxBackups: 2,
			MaxAge:     5,
			Compress:   true,
		},
		Warn: &logger.LoggerOptions{
			Filename:   "./logs/warn.log",
			MaxSize:    5,
			MaxBackups: 2,
			MaxAge:     5,
			Compress:   true,
		},
	}
	logger.Init(loggerOptions)
}
