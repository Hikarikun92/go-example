package rest

import (
	. "go-example/user"
	"testing"
)

func Test_facadeImpl_FindAll_withSuccess(t *testing.T) {
	entities := []*User{
		{Id: 1, Username: "Administrator"},
		{Id: 2, Username: "John Doe"},
		{Id: 3, Username: "Mary Doe"},
	}
	service := mockService{findAllImpl: func() []*User { return entities }}
	facade := facadeImpl{service: &service}

	dtos := facade.FindAll()
	if len(entities) != len(dtos) {
		t.Errorf("Expected %v DTOs, got %v", len(entities), len(dtos))
	}

	for i, dto := range dtos {
		expectedDto := ToReadDto(entities[i])
		if *dto != *expectedDto {
			t.Errorf("dtos[%v]: expected %v, got %v", i, expectedDto, dto)
		}
	}
}

type mockService struct {
	findAllImpl func() []*User
}

func (s *mockService) FindAll() []*User {
	return s.findAllImpl()
}
