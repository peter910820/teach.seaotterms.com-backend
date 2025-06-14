package dto

import "github.com/lib/pq"

type SeriesCreateResponse struct {
	Title string `gorm:"NOT NULL" json:"title"`
	Image string `json:"image"`
}

type SeriesModifyResponse struct {
	Title string `gorm:"NOT NULL" json:"title"`
	Image string `json:"image"`
}

type ArtilceCreateResponse struct {
	Title    string         `gorm:"NOT NULL" json:"title"`
	SeriesID uint           `gorm:"NOT NULL" json:"seriesId"`
	Image    string         `json:"image"`
	Tags     pq.StringArray `gorm:"type:text[]" json:"tags"`
	Content  string         `gorm:"NOT NULL" json:"content"`
}

type ArtilceModifyResponse struct {
	Title    string         `gorm:"NOT NULL" json:"title"`
	SeriesID uint           `gorm:"NOT NULL" json:"seriesId"`
	Image    string         `json:"image"`
	Tags     pq.StringArray `gorm:"type:text[]" json:"tags"`
	Content  string         `gorm:"NOT NULL" json:"content"`
}
