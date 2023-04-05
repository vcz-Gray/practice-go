package models

import (
	"time"
)

type User struct {
	Id            uint      `gorm:"primary_key" json:"id"`
	Username      string    `gorm:"not null" json:"username"`
	Email         string    `gorm:"not null" json:"email"`
	Password      string    `gorm:"not null" json:"password"`
	AuthGrade     int       `gorm:"default: 1" json:"authGrade"`
	BookmarkPosts string    `gorm:"default:'[]'" json:"bookmarkPosts"`
	LikeTags      string    `gorm:"default:'[]'" json:"likeTags"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
