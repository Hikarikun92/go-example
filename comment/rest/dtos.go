package rest

import (
	userRest "github.com/Hikarikun92/go-example/user/rest"
)

type CommentReadDto struct {
	Id            int                   `json:"id"`
	Title         string                `json:"title"`
	Body          string                `json:"body"`
	PublishedDate string                `json:"publishedDate"`
	User          *userRest.UserReadDto `json:"user"`
}
