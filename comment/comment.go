package comment

import (
	"go-example/user"
	"time"
)

type Comment struct {
	Id            int
	Title         string
	Body          string
	PublishedDate time.Time
	User          *user.User
}
