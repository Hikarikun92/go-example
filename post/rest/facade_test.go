package rest

import (
	"errors"
	"fmt"
	"github.com/Hikarikun92/go-example/comment"
	"github.com/Hikarikun92/go-example/post"
	"github.com/Hikarikun92/go-example/user"
	userRest "github.com/Hikarikun92/go-example/user/rest"
	"github.com/Hikarikun92/go-example/util"
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
			Comments: []*comment.Comment{
				{
					Id:            1,
					Title:         "Example comment 1",
					Body:          "Praesent sapien leo, viverra sed.",
					PublishedDate: time.Date(2021, time.January, 1, 18, 42, 32, 0, time.UTC),
					User:          &user.User{Id: 2, Username: "John Doe"},
				},
			},
		},
		{
			Id:            2,
			Title:         "Another example post",
			Body:          "Integer malesuada lorem non nunc.",
			PublishedDate: time.Date(2021, time.March, 15, 17, 53, 7, 0, time.UTC),
			User:          author,
			Comments:      []*comment.Comment{},
		},
	}

	service := mockService{findByUserIdImpl: func(userId int) ([]*post.Post, error) { return entities, nil }}

	facade := facadeImpl{service: &service}

	dtos, err := facade.FindByUserId(42)
	if err != nil {
		t.Error("Unexpected error", err)
	}

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
	findByUserIdImpl func(userId int) ([]*post.Post, error)
	findByIdImpl     func(id int) (*post.Post, error)
	createImpl       func(post *post.Post) (int, error)
}

func (s *mockService) FindByUserId(userId int) ([]*post.Post, error) {
	return s.findByUserIdImpl(userId)
}

func (s *mockService) FindById(id int) (*post.Post, error) {
	return s.findByIdImpl(id)
}

func (s *mockService) Create(post *post.Post) (int, error) {
	return s.createImpl(post)
}

func Test_facadeImpl_FindByUserId_withError(t *testing.T) {
	service := mockService{findByUserIdImpl: func(userId int) ([]*post.Post, error) { return nil, errors.New("Error finding posts") }}
	facade := facadeImpl{service: &service}

	dtos, err := facade.FindByUserId(42)
	if dtos != nil {
		t.Errorf("Call with error shouldn't return a value")
	}
	if err == nil {
		t.Error("Expected error, got none")
	}
}

func Test_facadeImpl_FindById_withSuccess(t *testing.T) {
	entity := &post.Post{
		Id:            33,
		Title:         "Example post",
		Body:          "Blabla",
		PublishedDate: time.Date(2022, time.January, 12, 15, 21, 25, 0, time.UTC),
		User:          &user.User{Id: 2, Username: "John Doe"},
		Comments: []*comment.Comment{
			{
				Id:            8,
				Title:         "Example comment 23",
				Body:          "Praesent sapien leo, viverra sed.",
				PublishedDate: time.Date(2022, time.January, 15, 18, 42, 32, 0, time.UTC),
				User:          &user.User{Id: 2, Username: "John Doe"},
			},
			{
				Id:            12,
				Title:         "Great article",
				Body:          "Nice example!",
				PublishedDate: time.Date(2022, time.February, 28, 7, 38, 12, 0, time.UTC),
				User:          &user.User{Id: 3, Username: "Mary Doe"},
			},
		},
	}
	service := mockService{findByIdImpl: func(id int) (*post.Post, error) { return entity, nil }}
	facade := facadeImpl{service: &service}

	//Due to the reference to *User, we can't build an expectedDto and compare it with !=
	dto, err := facade.FindById(entity.Id)
	if err != nil {
		t.Error("Unexpected error", err)
	}

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

	if userErrors := compareUserReadDtoWithEntity(dto.User, entity.User); len(userErrors) != 0 {
		for _, err := range userErrors {
			t.Errorf("dto.User." + err)
		}
	}

	if len(dto.Comments) != len(entity.Comments) {
		t.Errorf("Expected %v comments, got %v", len(entity.Comments), len(dto.Comments))
	}

	for i, commentDto := range dto.Comments {
		commentEntity := entity.Comments[i]

		if commentDto.Id != commentEntity.Id {
			t.Errorf("dto.Comments[%v].Id: expected %v, got %v", i, commentEntity.Id, commentDto.Id)
		}

		if commentDto.Title != commentEntity.Title {
			t.Errorf("dto.Comments[%v].Title: expected %v, got %v", i, commentEntity.Title, commentDto.Title)
		}

		if commentDto.Body != commentEntity.Body {
			t.Errorf("dto.Comments[%v].Body: expected %v, got %v", i, commentEntity.Body, commentDto.Body)
		}

		if userErrors := compareUserReadDtoWithEntity(commentDto.User, commentEntity.User); len(userErrors) != 0 {
			for _, err := range userErrors {
				t.Errorf("dto.Comments[%v].User."+err, i)
			}
		}
	}
}

func compareUserReadDtoWithEntity(dto *userRest.UserReadDto, entity *user.User) []string {
	result := []string{}
	if dto.Id != entity.Id {
		result = append(result, fmt.Sprintf("Id: expected %v, got %v", entity.Id, dto.Id))
	}

	if dto.Username != entity.Username {
		result = append(result, fmt.Sprintf("Username: expected %v, got %v", entity.Username, dto.Username))
	}

	return result
}

func Test_facadeImpl_FindById_withError(t *testing.T) {
	service := mockService{findByIdImpl: func(id int) (*post.Post, error) { return nil, errors.New("Error finding posts") }}
	facade := facadeImpl{service: &service}

	dto, err := facade.FindById(13)
	if dto != nil {
		t.Errorf("Call with error shouldn't return a value")
	}
	if err == nil {
		t.Error("Expected error, got none")
	}
}

func Test_facadeImpl_FindById_notFound(t *testing.T) {
	service := mockService{findByIdImpl: func(id int) (*post.Post, error) { return nil, nil }}
	facade := facadeImpl{service: &service}

	dto, err := facade.FindById(3)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if dto != nil {
		t.Errorf("Expected nil, got %v", dto)
	}
}
