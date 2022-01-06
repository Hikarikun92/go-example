package user

type Service interface {
	FindAll() []*User
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
