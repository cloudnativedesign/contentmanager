package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
	Author   Author `gorm:"embedded";embeddedPrefix:author_`
}

type CreateArticleInput struct {
	Title    string `json:"title" binding:"required"`
	Author   Author `gorm:"embedded";embeddedPrefix:author_`
	Subtitle string `json:"subtitle" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
