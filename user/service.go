package user

type Service interface {
	FindAll() ([]*User, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (s *serviceImpl) FindAll() ([]*User, error) {
	return s.repository.FindAll()
}
