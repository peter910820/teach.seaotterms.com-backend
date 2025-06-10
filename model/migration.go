package model

import (
	"os"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

func Migration(dbName string, db *gorm.DB) {
	switch dbName {
	case os.Getenv("DATABASE_NAME"):
		db.AutoMigrate(&Series{})
		db.AutoMigrate(&Article{})
		db.AutoMigrate(&Comment{})
	default:
		logrus.Fatal("error in migration function")
	}
}
