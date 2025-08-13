package connection

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Opts struct {
	Host, Port, User, Password, Name, SSLMode string
}

func New(opts Opts, isDev bool) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		opts.Host, opts.Port, opts.User, opts.Password, opts.Name, opts.SSLMode,
	)

	var gl logger.Interface = logger.Default.LogMode(logger.Silent)
	if isDev {
		gl = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 gl,
	})
	if err != nil {
		return nil, err
	}

	// connection pool (useful with managed PG)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	return db, nil
}
