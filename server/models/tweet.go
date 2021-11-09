package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Content  string `json:"content"`
	Author   Author `gorm:"embedded"`
	Reshares int32  `json:"reshares"`
}

type CreateTweetInput struct {
	Content string `json:"content"`
	Author  Author `json:"author";gorm:"embedded"`
}
