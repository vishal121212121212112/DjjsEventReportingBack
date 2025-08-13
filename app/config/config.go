package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DB struct {
	Host, Port, User, Password, Name, SSLMode string
	AutoMigrate bool
}

type Config struct {
	AppName   string
	AppEnv    string
	AppPort   string
	DB        DB
	JWTSecret string
}

func Load() *Config {
	// Decide environment (dev by default)
	env := os.Getenv("APP_ENV")
	if env == "" { env = "dev" }

	// Load .env.<env> into process env if present (safe to ignore if missing)
	_ = godotenv.Load(".env." + env)

	// Defaults (in case not set)
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("DB_AUTOMIGRATE", true)

	// Read from process env
	viper.AutomaticEnv()

	cfg := &Config{
		AppName:   viper.GetString("APP_NAME"),
		AppEnv:    env,
		AppPort:   viper.GetString("APP_PORT"),
		DB: DB{
			Host:        viper.GetString("DB_HOST"),
			Port:        viper.GetString("DB_PORT"),
			User:        viper.GetString("DB_USER"),
			Password:    viper.GetString("DB_PASSWORD"),
			Name:        viper.GetString("DB_NAME"),
			SSLMode:     viper.GetString("DB_SSLMODE"),
			AutoMigrate: viper.GetBool("DB_AUTOMIGRATE"),
		},
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
	log.Printf("config: loaded %s (.env.%s)", cfg.AppEnv, env)
	return cfg
}
