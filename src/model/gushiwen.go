package model

import "github.com/jinzhu/gorm"

// 朝代的表: 先秦、两汉、魏晋、南北朝、隋代、唐代、五代、宋代、金朝、元代、明代、清代
type Dynasty struct {
	gorm.Model
	Name string `gorm:"tye:varchar" json:"name"`
}

// 诗人的表
type Poet struct {
	gorm.Model
	Name        string `gorm:"type:varchar" json:"name"`
	ImageURL    string `gorm:"type:varchar" json:"image_url"`
	Number      uint   `gorm:"type:integer" json:"number"`
	Liked       uint   `gorm:"type:integer" json:"liked"`
	Description string `gorm:"type:varchar" json:"description"`
	DynastyID   uint   `gorm:"type:integer" json:"dynasty_id"`
	PoetryInfo  []PoetryInfo
}

// 诗类型的表: 诗、词、曲、文言文
type PoetryType struct {
	gorm.Model
	TypeName string `gorm:"type:varchar" json:"type_name"`
}

// 诗文的表
type PoetryInfo struct {
	gorm.Model
	Title     string `gorm:"type:varchar" json:"title"`
	Content   string `gorm:"type:varchar" json:"content"`
	Liked     uint   `gorm:"type:varchar;default(0)" json:"liked"`
	PoetID    uint   `gorm:"type:integer" json:"poet_id"`
	DynastyID uint   `gorm:"type:integer" json:"dynasty_id"`
}
