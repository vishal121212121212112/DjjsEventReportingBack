package registry

import (
	"event-reporting/app/api"
	"event-reporting/app/config"
	"event-reporting/app/database/pgsql/connection"
	"event-reporting/app/database/pgsql/repository"
	"event-reporting/app/handlers/user"
	"event-reporting/app/helpers/logger"
	"event-reporting/app/services/user"
)

func Build() *api.Routes {
	cfg := config.Load()
	log := logger.New(cfg.AppEnv)
	defer log.Sync()

	db, err := connection.New(connection.Opts{
		Host: cfg.DB.Host, Port: cfg.DB.Port, User: cfg.DB.User,
		Password: cfg.DB.Password, Name: cfg.DB.Name, SSLMode: cfg.DB.SSLMode,
	}, cfg.AppEnv == "dev")
	if err != nil { panic(err) }

	if cfg.DB.AutoMigrate {
		if err := connection.Bootstrap(db); err != nil { panic(err) }
	}

	userRepo := repository.NewUserRepository(db)
	userSvc  := services.NewUserService(userRepo)
	userH    := handlers.NewUserHandler(userSvc, cfg.JWTSecret)

	return &api.Routes{ User: userH, JWTSecret: cfg.JWTSecret }
}
