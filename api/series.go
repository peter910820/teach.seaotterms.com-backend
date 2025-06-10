package api

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"teach.seaotterms.com-backend/model"
)

// query series
func QuerySeries(c *fiber.Ctx, db *gorm.DB) error {
	var responseData []model.Series
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
		"msg":  "查詢Series資料成功",
		"data": responseData,
	})
}
