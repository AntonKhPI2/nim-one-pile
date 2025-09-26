package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/AntonKhPI2/nim-one-pile/internal/models"
)

func MustOpenMySQLGorm(ctx context.Context) *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		host := getenv("DB_HOST", "127.0.0.1")
		port := getenv("DB_PORT", "3306")
		user := getenv("DB_USER", "root")
		pass := os.Getenv("DB_PASSWORD")
		name := getenv("DB_NAME", "nim")
		params := getenv("DB_PARAMS", "parseTime=true&charset=utf8mb4&collation=utf8mb4_0900_ai_ci")

		if pass == "" {
			dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?%s", user, host, port, name, params)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pass, host, port, name, params)
		}
	}

	cfg := &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Warn),
	}
	db, err := gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		panic(fmt.Errorf("gorm open mysql: %w", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	if err := sqlDB.PingContext(ctx); err != nil {
		panic(fmt.Errorf("db ping: %w", err))
	}

	if err := db.WithContext(ctx).AutoMigrate(&models.Game{}); err != nil {
		panic(fmt.Errorf("auto-migrate: %w", err))
	}
	return db
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
