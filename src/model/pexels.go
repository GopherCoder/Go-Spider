package model

import "github.com/jinzhu/gorm"

type ImageSize struct {
	gorm.Model
	Width  string `gorm:"type:varchar" json:"width"`
	Height string `gorm:"type:varchar" json:"height"`
}

type ImageAddress struct {
	gorm.Model
	URL         string `gorm:"type:varchar" json:"url"`
	ImageSizeID uint   `grom:"type:integer" json:"image_size_id"`
	SizeType    string `gorm:"type:varchar" json:"size_type"`
}
