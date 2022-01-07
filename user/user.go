package user

import "strconv"

type User struct {
	Id       int
	Username string
}

func (u *User) String() string {
	return "User(Id=" + strconv.Itoa(u.Id) + ", Username=" + u.Username + ")"
}

type Credentials struct {
	*User
	Password string
	Roles    []string
}
