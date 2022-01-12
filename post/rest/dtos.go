package rest

import (
	user_rest "go-example/user/rest"
)

type PostByUserDto struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	PublishedDate string `json:"publishedDate"`
}

type PostByIdDto struct {
	Id            int                    `json:"id"`
	Title         string                 `json:"title"`
	Body          string                 `json:"body"`
	PublishedDate string                 `json:"publishedDate"`
	User          *user_rest.UserReadDto `json:"user"`
}
