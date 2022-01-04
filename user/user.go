package user

type User struct {
	Id       int
	Username string
}

type Credentials struct {
	User
	Password string
	Roles    []string
}
