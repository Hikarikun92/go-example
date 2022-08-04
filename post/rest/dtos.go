package rest

import (
	commentRest "github.com/Hikarikun92/go-example/comment/rest"
	userRest "github.com/Hikarikun92/go-example/user/rest"
)

type PostByUserDto struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	PublishedDate string `json:"publishedDate"`
}

type PostByIdDto struct {
	Id            int                           `json:"id"`
	Title         string                        `json:"title"`
	Body          string                        `json:"body"`
	PublishedDate string                        `json:"publishedDate"`
	User          *userRest.UserReadDto         `json:"user"`
	Comments      []*commentRest.CommentReadDto `json:"comments"`
}

type CreatePostDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
