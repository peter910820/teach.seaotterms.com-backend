package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"teach.seaotterms.com-backend/api"
)

func ApiRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	apiGroup.Get("/series", func(c *fiber.Ctx) error {
		return api.QuerySeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})
}
