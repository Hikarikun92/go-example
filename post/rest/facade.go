package rest

import (
	"github.com/Hikarikun92/go-example/comment/rest"
	"github.com/Hikarikun92/go-example/post"
	"github.com/Hikarikun92/go-example/user"
	userRest "github.com/Hikarikun92/go-example/user/rest"
	"github.com/Hikarikun92/go-example/util"
	"time"
)

type Facade interface {
	FindByUserId(userId int) ([]*PostByUserDto, error)
	FindById(id int) (*PostByIdDto, error)
	Create(dto *CreatePostDto, userId int) (int, error)
}

func NewFacade(service post.Service) Facade {
	return &facadeImpl{service: service}
}

type facadeImpl struct {
	service post.Service
}

func (f *facadeImpl) FindByUserId(userId int) ([]*PostByUserDto, error) {
	entities, err := f.service.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	dtos := make([]*PostByUserDto, len(entities))

	for i, entity := range entities {
		dtos[i] = toPostByUserDto(entity)
	}

	return dtos, nil
}

func toPostByUserDto(post *post.Post) *PostByUserDto {
	return &PostByUserDto{
		Id:            post.Id,
		Title:         post.Title,
		Body:          post.Body,
		PublishedDate: util.TimeToIso(post.PublishedDate),
	}
}

func (f *facadeImpl) FindById(id int) (*PostByIdDto, error) {
	p, err := f.service.FindById(id)
	if err != nil {
		return nil, err
	}
	return toPostByIdDto(p), nil
}

func toPostByIdDto(post *post.Post) *PostByIdDto {
	if post == nil {
		return nil
	}

	//Cache to avoid creating DTOs of already converted users
	userCache := make(map[int]*userRest.UserReadDto)

	comments := make([]*rest.CommentReadDto, len(post.Comments))
	for i, comment := range post.Comments {
		userDto, ok := userCache[comment.User.Id]
		if !ok {
			userDto = userRest.ToReadDto(comment.User)
			userCache[comment.User.Id] = userDto
		}

		comments[i] = &rest.CommentReadDto{
			Id:            comment.Id,
			Title:         comment.Title,
			Body:          comment.Body,
			PublishedDate: util.TimeToIso(comment.PublishedDate),
			User:          userDto,
		}
	}

	userDto, ok := userCache[post.User.Id]
	if !ok {
		userDto = userRest.ToReadDto(post.User)
		userCache[post.User.Id] = userDto
	}

	return &PostByIdDto{
		Id:            post.Id,
		Title:         post.Title,
		Body:          post.Body,
		PublishedDate: util.TimeToIso(post.PublishedDate),
		User:          userDto,
		Comments:      comments,
	}
}

func (f *facadeImpl) Create(dto *CreatePostDto, userId int) (int, error) {
	p := &post.Post{
		Title:         dto.Title,
		Body:          dto.Body,
		PublishedDate: time.Now(),
		User:          &user.User{Id: userId},
	}

	return f.service.Create(p)
}
