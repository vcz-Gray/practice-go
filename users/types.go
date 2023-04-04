package users

import "time"

type User struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ChangedPassword string `json:"changedPassword"`
}

type Article struct {
	Id          int
	UserId      int
	Title       string
	Description string
	Date        time.Time
}
