package model

import "github.com/jinzhu/gorm"

type Repositories struct {
	gorm.Model
	Name        string `gorm:"type:varchar" json:"name"`
	TotalStar   string `gorm:"type:varchar" json:"total_star"`
	SinceStar   string `gorm:"type:varchar" json:"since_star"`
	Fork        string `gorm:"type:varchar" json:"fork"`
	Description string `gorm:"type:varchar" json:"description"`
	SinceFlag   string `gorm:"type:varchar" json:"since_flag"`
	Language    string `gorm:"type:varchar" json:"language"`
}
