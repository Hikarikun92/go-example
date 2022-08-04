package rest

import (
	"github.com/Hikarikun92/go-example/user"
	"strconv"
)

type UserReadDto struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u *UserReadDto) String() string {
	return "UserReadDto(Id=" + strconv.Itoa(u.Id) + ", Username=" + u.Username + ")"
}

func ToReadDto(u *user.User) *UserReadDto {
	return &UserReadDto{Id: u.Id, Username: u.Username}
}
