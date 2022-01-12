package rest

import "go-example/user"

type Facade interface {
	FindAll() []*UserReadDto
}

func NewFacade(service user.Service) Facade {
	return &facadeImpl{service: service}
}

type facadeImpl struct {
	service user.Service
}

func (f *facadeImpl) FindAll() []*UserReadDto {
	return ToReadDtos(f.service.FindAll())
}
