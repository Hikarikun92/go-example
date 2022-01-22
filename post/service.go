package post

type Service interface {
	FindByUserId(userId int) ([]*Post, error)
	FindById(id int) (*Post, error)
	Create(post *Post) (int, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (s *serviceImpl) FindByUserId(userId int) ([]*Post, error) {
	return s.repository.FindByUserId(userId)
}

func (s *serviceImpl) FindById(id int) (*Post, error) {
	return s.repository.FindById(id)
}

func (s *serviceImpl) Create(post *Post) (int, error) {
	return s.repository.Create(post)
}
