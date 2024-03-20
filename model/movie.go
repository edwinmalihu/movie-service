package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Rating      float64
	Image       string `gorm:"type:varchar(255)"`
}

func (Movie) TableName() string {
	return "movie"
}
