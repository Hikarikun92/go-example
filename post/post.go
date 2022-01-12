package post

import (
	"go-example/user"
	"time"
)

type Post struct {
	Id            int
	Title         string
	Body          string
	PublishedDate time.Time
	User          *user.User
}
