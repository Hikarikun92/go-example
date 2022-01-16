package rest

import "go-example/user"

type Facade interface {
	FindAll() ([]*UserReadDto, error)
}

func NewFacade(service user.Service) Facade {
	return &facadeImpl{service: service}
}

type facadeImpl struct {
	service user.Service
}

func (f *facadeImpl) FindAll() ([]*UserReadDto, error) {
	users, err := f.service.FindAll()
	if err != nil {
		return nil, err
	}
	return ToReadDtos(users), nil
}
