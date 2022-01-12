package post

type Service interface {
	FindByUserId(userId int) []*Post
	FindById(id int) *Post
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (s *serviceImpl) FindByUserId(userId int) []*Post {
	return s.repository.FindByUserId(userId)
}

func (s *serviceImpl) FindById(id int) *Post {
	return s.repository.FindById(id)
}
