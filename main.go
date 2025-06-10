package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"teach.seaotterms.com-backend/model"
	"teach.seaotterms.com-backend/router"
)

var (
	// management database connect
	dbs = make(map[string]*gorm.DB)
)

func init() {
	// init logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	// init env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file load error: %v", err)
	}
}

func main() {
	// init migration
	dbName, db := model.InitDsn(os.Getenv("DATABASE_NAME"))
	dbs[dbName] = db
	model.Migration(dbName, dbs[dbName])

	app := fiber.New()

	// route group
	apiGroup := app.Group("/api") // main api route group
	router.ApiRouter(apiGroup, dbs)

	logrus.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", os.Getenv("PRODUCTION_PORT"))))
}
