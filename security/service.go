package security

import (
	"errors"
	"go-example/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(username string, password string) (string, error)
}

type serviceImpl struct {
	repository user.Repository
	jwtService JwtService
}

func NewService(repository user.Repository, jwtService JwtService) Service {
	return &serviceImpl{repository: repository, jwtService: jwtService}
}

func (s *serviceImpl) Login(username string, password string) (string, error) {
	credentials, err := s.repository.FindCredentialsByUsername(username)
	if credentials == nil {
		return "", errors.New("Unknown user: " + username)
	}
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := s.jwtService.GenerateToken(username)
	return token, nil
}
