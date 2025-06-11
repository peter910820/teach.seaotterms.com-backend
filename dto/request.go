package dto

import "github.com/lib/pq"

type SeriesCreateResponse struct {
	Title string `gorm:"NOT NULL" json:"title"`
}

type SeriesModifyResponse struct {
	Title string `gorm:"NOT NULL" json:"title"`
}

type ArtilceCreateResponse struct {
	Title    string         `gorm:"NOT NULL" json:"title"`
	SeriesID uint           `gorm:"NOT NULL" json:"seriesId"`
	Tags     pq.StringArray `gorm:"type:text[]" json:"tags"`
	Content  string         `gorm:"NOT NULL" json:"content"`
}
