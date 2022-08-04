package rest

import (
	. "github.com/Hikarikun92/go-example/user"
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
