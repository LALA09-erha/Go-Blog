package models

import "time"

type User struct {
	UserID           string    `gorm:"primaryKey"`
	Username         string    `gorm:"unique;not null"`
	Email            string    `gorm:"unique;not null"`
	PasswordHash     string    `gorm:"not null"`
	DisplayName      string    `gorm:"not null"`
	Bio              string    `gorm:"type:text"`
	ProfilePictureURL string   `gorm:"default:'#'"`
	Role             string    `gorm:"default:'USER'"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Category struct {
	CategoryID string `gorm:"primaryKey"`
	Name       string `gorm:"unique;not null"`
	Slug       string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Article struct {
	ArticleID  string    `gorm:"primaryKey"`
	CategoryID string    `gorm:"not null"`
	Title      string    `gorm:"not null"`
	Slug       string    `gorm:"not null"`
	Content    string    `gorm:"type:text;not null"`
	AuthorID   string    `gorm:"not null"`
	ViewCount  int       `gorm:"default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
