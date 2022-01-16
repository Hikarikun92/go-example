package rest

import (
	"fmt"
	"go-example/comment"
	"go-example/post"
	"go-example/user"
	userRest "go-example/user/rest"
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

func Test_facadeImpl_FindById_notFound(t *testing.T) {
	service := mockService{findByIdImpl: func(id int) *post.Post { return nil }}
	facade := facadeImpl{service: &service}

	dto := facade.FindById(3)
	if dto != nil {
		t.Errorf("Expected nil, got %v", dto)
	}
}
