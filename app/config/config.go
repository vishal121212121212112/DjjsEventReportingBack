package config

import (
	"log"
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
	viper.SetConfigFile(".env.dev") // or select by APP_ENV
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config: no file read (%v); relying on env only", err)
	}
	return &Config{
		AppName: viper.GetString("APP_NAME"),
		AppEnv:  viper.GetString("APP_ENV"),
		AppPort: viper.GetString("APP_PORT"),
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
}
