package models

import "time"

type Person struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	DateJoined   time.Time `json:"date_joined"`
	UserName     string    `json:"user_name"`
	EmailPrimary string    `json:"primary_email"`
}

type Author struct {
	Person    Person `gorm:"embedded"`
	Articles  []Article
	Followers []Person
}
