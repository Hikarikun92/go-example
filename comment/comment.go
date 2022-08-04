package comment

import (
	"github.com/Hikarikun92/go-example/user"
	"time"
)

type Comment struct {
	Id            int
	Title         string
	Body          string
	PublishedDate time.Time
	User          *user.User
}
