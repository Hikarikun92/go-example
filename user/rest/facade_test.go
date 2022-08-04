package rest

import (
	"errors"
	. "github.com/Hikarikun92/go-example/user"
	"testing"
)

func Test_facadeImpl_FindAll_withSuccess(t *testing.T) {
	entities := []*User{
		{Id: 1, Username: "Administrator"},
		{Id: 2, Username: "John Doe"},
		{Id: 3, Username: "Mary Doe"},
	}
	service := mockService{findAllImpl: func() ([]*User, error) { return entities, nil }}
	facade := facadeImpl{service: &service}

	dtos, err := facade.FindAll()
	if err != nil {
		t.Error("Unexpected error", err)
	}

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
	findAllImpl func() ([]*User, error)
	loginImpl   func(username string, password string) (string, error)
}

func (s *mockService) FindAll() ([]*User, error) {
	return s.findAllImpl()
}

func (s *mockService) Login(username string, password string) (string, error) {
	return s.loginImpl(username, password)
}

func Test_facadeImpl_FindAll_withError(t *testing.T) {
	service := mockService{findAllImpl: func() ([]*User, error) { return nil, errors.New("Error finding users") }}
	facade := facadeImpl{service: &service}

	dtos, err := facade.FindAll()
	if dtos != nil {
		t.Errorf("Call with error shouldn't return a value")
	}
	if err == nil {
		t.Error("Expected error, got none")
	}
}
