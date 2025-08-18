package connection

import (
	"database/sql"
	"event-reporting/app/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	if err := ensurePgcrypto(db); err != nil {
		log.Fatalf("❌ Failed to ensure pgcrypto: %v", err)
	}
	log.Println("🔒 pgcrypto extension is ready (gen_random_uuid available)")

	// Run Migrations
	log.Println("🔄 Running database migrations...")
	err = db.AutoMigrate(
		&models.User{},
		&models.EventHistory{},
		&models.GuestMaster{},
		&models.MediaAndDocumentation{},
		&models.ProgramDonation{},
		&models.ProgramVolunteer{},
		&models.ProgramMaster{},

		// geographies
		&models.Country{},
		&models.State{},
		&models.District{},
		&models.City{},
	)

	if err != nil {
		log.Fatal("❌ Failed to migrate models:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	// Assign connection to global variable
	Db = db
	log.Println("🚀 Database connection initialized successfully!")
}

func ensurePgcrypto(db *gorm.DB) error {
	// CREATE EXTENSION is idempotent
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto`).Error; err != nil {
		return fmt.Errorf("create extension pgcrypto: %w", err)
	}
	// optional: quick sanity check that function is callable
	var dummy sql.NullString
	if err := db.Raw(`SELECT gen_random_uuid()::text`).Scan(&dummy).Error; err != nil {
		return fmt.Errorf("gen_random_uuid not available (pgcrypto not loaded?): %w", err)
	}
	return nil
}
