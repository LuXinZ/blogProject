package blogApp

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Content string `gorm:"type:text;not null"`
	Author  string `gorm:"varchar(36)"`
}
