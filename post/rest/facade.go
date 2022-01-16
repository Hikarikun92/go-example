package rest

import (
	"go-example/comment/rest"
	"go-example/post"
	userRest "go-example/user/rest"
	"go-example/util"
)

type Facade interface {
	FindByUserId(userId int) []*PostByUserDto
	FindById(id int) *PostByIdDto
}

func NewFacade(service post.Service) Facade {
	return &facadeImpl{service: service}
}

type facadeImpl struct {
	service post.Service
}

func (f *facadeImpl) FindByUserId(userId int) []*PostByUserDto {
	entities := f.service.FindByUserId(userId)
	dtos := make([]*PostByUserDto, len(entities))

	for i, entity := range entities {
		dtos[i] = toPostByUserDto(entity)
	}

	return dtos
}

func toPostByUserDto(post *post.Post) *PostByUserDto {
	return &PostByUserDto{
		Id:            post.Id,
		Title:         post.Title,
		Body:          post.Body,
		PublishedDate: util.TimeToIso(post.PublishedDate),
	}
}

func (f *facadeImpl) FindById(id int) *PostByIdDto {
	return toPostByIdDto(f.service.FindById(id))
}

func toPostByIdDto(post *post.Post) *PostByIdDto {
	if post == nil {
		return nil
	}

	comments := make([]*rest.CommentReadDto, len(post.Comments))
	for i, comment := range post.Comments {
		comments[i] = &rest.CommentReadDto{
			Id:            comment.Id,
			Title:         comment.Title,
			Body:          comment.Body,
			PublishedDate: util.TimeToIso(comment.PublishedDate),
			User:          userRest.ToReadDto(comment.User),
		}
	}

	return &PostByIdDto{
		Id:            post.Id,
		Title:         post.Title,
		Body:          post.Body,
		PublishedDate: util.TimeToIso(post.PublishedDate),
		User:          userRest.ToReadDto(post.User),
		Comments:      comments,
	}
}
