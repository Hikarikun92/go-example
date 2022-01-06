package post

import (
	"go-example/user"
	"time"
)

type Post struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	PublishedDate time.Time `json:"publishedDate"`
	User          *user.User
}
