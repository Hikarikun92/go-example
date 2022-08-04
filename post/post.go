package post

import (
	"github.com/Hikarikun92/go-example/comment"
	"github.com/Hikarikun92/go-example/user"
	"time"
)

type Post struct {
	Id            int
	Title         string
	Body          string
	PublishedDate time.Time
	User          *user.User
	Comments      []*comment.Comment
}
