package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"teach.seaotterms.com-backend/api"
)

func SeriesRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	apiGroup.Get("/series", func(c *fiber.Ctx) error {
		return api.QuerySeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Post("/series", func(c *fiber.Ctx) error {
		return api.CreateSeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Patch("/series/:id", func(c *fiber.Ctx) error {
		return api.ModifySeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})
}

func ArticleApiRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
}

func CommentApiRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
}
