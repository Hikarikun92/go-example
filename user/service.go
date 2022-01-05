package user

type Service interface {
	FindAll() []*User
	FindCredentialsByUsername(username string) *Credentials
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (s *serviceImpl) FindAll() []*User {
	return s.repository.FindAll()
}

func (s *serviceImpl) FindCredentialsByUsername(username string) *Credentials {
	return s.repository.FindCredentialsByUsername(username)
}
