package model

import (
	"time"

	"github.com/lib/pq"
)

type Series struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"NOT NULL" json:"title"`
	Image         string    `json:"image"`
	ArticleAmount uint      `gorm:"NOT NULL; default:0" json:"articleAmount"`
	CreatedAt     time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName   string    `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedName   string    `json:"updatedName"`
}

type Article struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"NOT NULL" json:"title"`
	Image       string         `json:"image"`
	SeriesID    uint           `gorm:"NOT NULL" json:"seriesId"` // series table id
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	Content     string         `gorm:"NOT NULL" json:"content"`
	CreatedAt   time.Time      `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName string         `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	UpdatedName string         `json:"updatedName"`
}

type Comment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Content     string    `gorm:"NOT NULL" json:"content"`
	ArticleID   uint      `gorm:"NOT NULL" json:"articleId"` // article table id
	CreatedAt   time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName string    `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedName string    `json:"updatedName"`
}
