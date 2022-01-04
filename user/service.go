package user

type Repository interface {
	FindAll() []User
	FindCredentialsByUsername(username string) *Credentials
}

type Service struct {
	Repository
}

func (s Service) FindAll() []User {
	return s.Repository.FindAll()
}

func (s Service) FindCredentialsByUsername(username string) *Credentials {
	return s.Repository.FindCredentialsByUsername(username)
}
