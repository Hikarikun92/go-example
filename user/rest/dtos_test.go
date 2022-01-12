package rest

import (
	. "go-example/user"
	"testing"
)

func TestToReadDto(t *testing.T) {
	entity := User{Id: 3, Username: "Some user"}
	dto := *ToReadDto(&entity)

	if dto.Id != entity.Id {
		t.Errorf("dto.Id: expected %v, got %v", entity.Id, dto.Id)
	}

	if dto.Username != entity.Username {
		t.Errorf("dto.Username: expected %v, got %v", entity.Username, dto.Username)
	}
}

func TestToReadDtos(t *testing.T) {
	entities := []*User{
		{Id: 1, Username: "User 1"},
		{Id: 2, Username: "User 2"},
	}
	dtos := ToReadDtos(entities)
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
