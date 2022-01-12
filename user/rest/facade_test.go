package rest

import (
	. "go-example/user"
	"testing"
)

func Test_facadeImpl_FindAll_withSuccess(t *testing.T) {
	service := mockService{findAllImpl: func() []*User {
		return []*User{
			{Id: 1, Username: "Administrator"},
			{Id: 2, Username: "John Doe"},
			{Id: 3, Username: "Mary Doe"},
		}
	}}
	facade := facadeImpl{service: service}

	entities := service.FindAll()
	dtos := facade.FindAll()
	if len(entities) != len(dtos) {
		t.Errorf("Expected %v DTOs, got %v", len(entities), len(dtos))
	}

	for i, dto := range dtos {
		entity := entities[i]
		if dto.Id != entity.Id {
			t.Errorf("dtos[%v].Id: expected %v, got %v", i, entity.Id, dto.Id)
		}

		if dto.Username != entity.Username {
			t.Errorf("dtos[%v].Username: expected %v, got %v", i, entity.Username, dto.Username)
		}
	}
}

type mockService struct {
	findAllImpl func() []*User
}

func (s mockService) FindAll() []*User {
	return s.findAllImpl()
}
