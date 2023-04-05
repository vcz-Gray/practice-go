package interfaces

import "time"

type User interface {
	ID() int
	USERNAME() string
	EMAIL() string
	PASSWORD() string
	AUTHGRADE() int
	BOOKMARKPOSTS() string
	LIKETAGS() string
	CREATEDAT() time.Time
	UPDATEDAT() time.Time
}
