package rest

import (
	"go-example/post"
	"go-example/user"
	"go-example/util"
	"testing"
	"time"
)

func Test_facadeImpl_FindByUserId_withSuccess(t *testing.T) {
	author := &user.User{Id: 42, Username: "Username"}
	entities := []*post.Post{
		{
			Id:            1,
			Title:         "Example post no. 1",
			Body:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse placerat.",
			PublishedDate: time.Date(2021, time.January, 1, 12, 3, 18, 0, time.UTC),
			User:          author,
		},
		{
			Id:            2,
			Title:         "Another example post",
			Body:          "Integer malesuada lorem non nunc.",
			PublishedDate: time.Date(2021, time.March, 15, 17, 53, 7, 0, time.UTC),
			User:          author,
		},
	}

	service := mockService{findByUserIdImpl: func(userId int) []*post.Post { return entities }}

	facade := facadeImpl{service: &service}

	dtos := facade.FindByUserId(42)
	if len(dtos) != len(entities) {
		t.Errorf("Expected %v DTOs, got %v", len(entities), len(dtos))
	}

	for i, dto := range dtos {
		entity := entities[i]
		expectedDto := PostByUserDto{Id: entity.Id, Title: entity.Title, Body: entity.Body, PublishedDate: util.TimeToIso(entity.PublishedDate)}

		if *dto != expectedDto {
			t.Errorf("dtos[%v]: expected %v, got %v", i, expectedDto, dto)
		}
	}
}

type mockService struct {
	findByUserIdImpl func(userId int) []*post.Post
	findByIdImpl     func(id int) *post.Post
}

func (s *mockService) FindByUserId(userId int) []*post.Post {
	return s.findByUserIdImpl(userId)
}

func (s *mockService) FindById(id int) *post.Post {
	return s.findByIdImpl(id)
}

func Test_facadeImpl_FindById_withSuccess(t *testing.T) {
	entity := &post.Post{
		Id:            33,
		Title:         "Example post",
		Body:          "Blabla",
		PublishedDate: time.Date(2022, time.January, 12, 15, 21, 25, 0, time.UTC),
		User:          &user.User{Id: 2, Username: "Example user"},
	}
	service := mockService{findByIdImpl: func(id int) *post.Post { return entity }}
	facade := facadeImpl{service: &service}

	//Due to the reference to *User, we can't build an expectedDto and compare it with !=
	dto := facade.FindById(entity.Id)
	if dto.Id != entity.Id {
		t.Errorf("dto.Id: expected %v, got %v", entity.Id, dto.Id)
	}
	if dto.Title != entity.Title {
		t.Errorf("dto.Title: expected %v, got %v", entity.Title, dto.Title)
	}
	if dto.Body != entity.Body {
		t.Errorf("dto.Body: expected %v, got %v", entity.Body, dto.Body)
	}

	expectedPublishedDate := util.TimeToIso(entity.PublishedDate)
	if dto.PublishedDate != expectedPublishedDate {
		t.Errorf("dto.PublishedDate: expected %v, got %v", expectedPublishedDate, dto.PublishedDate)
	}
}

func Test_facadeImpl_FindById_notFound(t *testing.T) {
	service := mockService{findByIdImpl: func(id int) *post.Post { return nil }}
	facade := facadeImpl{service: &service}

	dto := facade.FindById(3)
	if dto != nil {
		t.Errorf("Expected nil, got %v", dto)
	}
}
