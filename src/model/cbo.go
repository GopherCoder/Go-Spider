package model

import "github.com/jinzhu/gorm"

type RankListCBOInfo struct {
	gorm.Model
	BoxOffice    string `gorm:"type:varchar"json:"box_office"`
	MovieID      string `gorm:"type:varchar"json:"movie_id"`
	MovieEnName  string `gorm:"type:varchar"json:"movie_en_name"`
	MovieName    string `gorm:"type:varchar"json:"movie_name"`
	Ranking      string `gorm:"type:varchar"json:"ranking"`
	DefaultImage string `gorm:"type:varchar"json:"default_image"`
	ReleaseYear  string `gorm:"type:varchar"json:"release_year"`
	Country      string `gorm:"type:varchar"json:"country"`
}

func (RankListCBOInfo) TableName() string {
	return "cbo_movies"
}
