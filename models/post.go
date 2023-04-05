package models

import "time"

type Post struct {
	Id        int       `json:"id"`
	AuthorId  string    `json:"authorId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
