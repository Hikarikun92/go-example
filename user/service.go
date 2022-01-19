package user

import (
	"go-example/security"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]*User, error)
	Login(username string, password string) (string, error)
}

type serviceImpl struct {
	repository            Repository
	authenticationManager security.AuthenticationManager
}

func NewService(repository Repository, authenticationManager security.AuthenticationManager) Service {
	return &serviceImpl{repository: repository, authenticationManager: authenticationManager}
}

func (s *serviceImpl) FindAll() ([]*User, error) {
	return s.repository.FindAll()
}

func (s *serviceImpl) Login(username string, password string) (string, error) {
	credentials, err := s.repository.FindCredentialsByUsername(username)
	if credentials == nil || err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := s.authenticationManager.GenerateToken(username)
	return token, nil
}
