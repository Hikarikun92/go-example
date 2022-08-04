package rest

import "github.com/Hikarikun92/go-example/security"

type Facade interface {
	Login(dto *LoginDto) (string, error)
}

type facadeImpl struct {
	service security.Service
}

func NewFacade(service security.Service) Facade {
	return &facadeImpl{service: service}
}

func (f *facadeImpl) Login(dto *LoginDto) (string, error) {
	return f.service.Login(dto.Username, dto.Password)
}
