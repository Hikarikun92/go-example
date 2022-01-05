package user

type Repository interface {
	FindAll() []*User
	FindCredentialsByUsername(username string) *Credentials
}

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) FindAll() []*User {
	return []*User{
		{Id: 1, Username: "user1"},
		{Id: 2, Username: "user2"},
	}
}

func (r *repositoryImpl) FindCredentialsByUsername(username string) *Credentials {
	switch username {
	case "user1":
		return &Credentials{User: &User{Id: 1, Username: "user1"}, Password: "pass1", Roles: []string{"ROLE_ADMIN", "ROLE_USER"}}
	case "user2":
		return &Credentials{User: &User{Id: 2, Username: "user2"}, Password: "pass2", Roles: []string{"ROLE_USER"}}
	default:
		return nil
	}
}
