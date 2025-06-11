package api

import (
	"errors"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"teach.seaotterms.com-backend/dto"
	"teach.seaotterms.com-backend/model"
)

type LinkUpdateSeries struct {
	ArticleAmount uint
	UpdateTime    time.Time
}

// query series
func QueryArticle(c *fiber.Ctx, db *gorm.DB) error {
	var responseData []model.Article
	var r *gorm.DB
	// URL decoding
	id, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if id == "" {
		r = db.Order("COALESCE(updated_at, created_at) DESC").Find(&responseData)
	} else {
		r = db.Where("id = ?", id).Order("COALESCE(updated_at, created_at) DESC").Find(&responseData)
	}
	if r.Error != nil {
		// if record not exist
		if r.Error == gorm.ErrRecordNotFound {
			logrus.Error(r.Error)
			//404
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": r.Error.Error(),
			})
		} else {
			logrus.Fatal(r.Error.Error())
			// 500
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": r.Error.Error(),
			})
		}
	}
	logrus.Info("Query series table success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "查詢Article資料成功",
		"data": responseData,
	})
}

func CreateArticle(c *fiber.Ctx, db *gorm.DB) error {
	var clientData dto.ArtilceCreateResponse
	// load client data
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	data := model.Article{
		Title:       clientData.Title,
		SeriesID:    clientData.SeriesID,
		Tags:        clientData.Tags,
		Content:     clientData.Content,
		CreatedAt:   time.Now(),
		CreatedName: "Root",
	}

	// confirm whether a exists
	var seriesData model.Series
	r := db.First(&seriesData, clientData.SeriesID)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "資料不存在",
		})
	} else if r.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": r.Error,
		})
	}

	r = db.Create(&data)
	if r.Error != nil {
		logrus.Error(r.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": r.Error.Error(),
		})
	}

	r = db.Model(&model.Series{}).Where("id = ?", clientData.SeriesID).
		Select("article_amount", "updated_at").
		Updates(LinkUpdateSeries{
			ArticleAmount: seriesData.ArticleAmount + 1,
			UpdateTime:    time.Now(),
		})
	if r.Error != nil {
		logrus.Error(r.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": r.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "新增Article資料成功",
		"data": data,
	})
}
