package model

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDsn(dbName string) (string, *gorm.DB) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_OWNER"),
		os.Getenv("DATABASE_PASSWORD"),
		dbName,
		os.Getenv("DATABASE_PORT"))

	// get connect db variable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("connect database failed: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("cannot get sql.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return dbName, db
}
