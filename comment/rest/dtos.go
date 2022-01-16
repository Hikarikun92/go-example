package rest

import (
	userRest "go-example/user/rest"
)

type CommentReadDto struct {
	Id            int                   `json:"id"`
	Title         string                `json:"title"`
	Body          string                `json:"body"`
	PublishedDate string                `json:"publishedDate"`
	User          *userRest.UserReadDto `json:"user"`
}
