package rest

import "github.com/Hikarikun92/go-example/user"

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

	dtos := make([]*UserReadDto, len(users))
	for i, u := range users {
		dtos[i] = ToReadDto(u)
	}
	return dtos, nil
}
