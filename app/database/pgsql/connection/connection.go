package connection

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// pgmodels "event-reporting/app/pgModels"
)

var Db *gorm.DB // Global variable to hold DB connection

// ConnectDatabase initializes the PostgreSQL connection using GORM
func ConnectDatabase() {
	log.Println("📢 Starting database connection...")

	// Load .env file
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	log.Println("✅ .env file loaded successfully")

	// Read database credentials from .env
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	// Create PostgreSQL DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	// Create a new GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Output to console
		logger.Config{
			SlowThreshold:             time.Second, // Log queries slower than 1 second
			LogLevel:                  logger.Info, // Log all SQL queries
			IgnoreRecordNotFoundError: true,        // Ignore "record not found" errors
			Colorful:                  true,        // Enable colorful logs
		},
	)

	log.Println("🔗 Connecting to PostgreSQL database...")

	// Open database connection with query logging enabled
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatalf("❌ Failed to connect to Postgres DB: %v", err)
	}
	log.Println("✅ Successfully connected to PostgreSQL!")

	// Run Migrations
	log.Println("🔄 Running database migrations...")
	// err = db.AutoMigrate(
	// 	&pgmodels.AssignedContacts{},
	// 	&pgmodels.RequestedChats{},
	// 	&pgmodels.ImportContactReport{},
	// 	&pgmodels.AgentReport{},
	// 	&pgmodels.FollowUp{},
	// 	&pgmodels.FollowUpRemarks{},
	// )
	if err != nil {
		log.Fatal("❌ Failed to migrate models:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	// Assign connection to global variable
	Db = db
	log.Println("🚀 Database connection initialized successfully!")
}
