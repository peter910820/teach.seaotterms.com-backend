package dto

type SeriesCreateResponse struct {
	Title string `gorm:"NOT NULL" json:"title"`
}

type SeriesModifyResponse struct {
	Title string `gorm:"NOT NULL" json:"title"`
}
