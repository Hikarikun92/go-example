package user

import "strconv"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u *User) String() string {
	return "User(Id=" + strconv.Itoa(u.Id) + ", Username=" + u.Username + ")"
}

type Credentials struct {
	*User
	Password string
	Roles    []string
}
