package security

import (
	"errors"
	"github.com/gorilla/mux"
	u "github.com/Hikarikun92/go-example/user"
	"testing"
)

func Test_serviceImpl_Login_withSuccess(t *testing.T) {
	user := &u.User{Id: 2, Username: "John Doe"}
	credentials := &u.Credentials{
		User:     user,
		Password: "$2a$10$bS.HuGI.l5pFEjfjDIjB2.3t9h62kRSi3exUTBhbs6vqrJouNTDh2",
		Roles:    []string{"ROLE_USER"},
	}

	repository := &mockRepository{findCredentialsByUsernameImpl: func(username string) (*u.Credentials, error) { return credentials, nil }}
	jwtService := &mockJwtService{generateTokenImpl: func(username string) string { return "example-token" }}

	service := serviceImpl{repository: repository, jwtService: jwtService}
	token, err := service.Login(user.Username, "PaSsW0rD!")
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if token != "example-token" {
		t.Errorf("Expected example-token, got %v", token)
	}
}

type mockRepository struct {
	findAllImpl                   func() ([]*u.User, error)
	findCredentialsByUsernameImpl func(username string) (*u.Credentials, error)
}

func (r *mockRepository) FindAll() ([]*u.User, error) {
	return r.findAllImpl()
}

func (r *mockRepository) FindCredentialsByUsername(username string) (*u.Credentials, error) {
	return r.findCredentialsByUsernameImpl(username)
}

type mockJwtService struct {
	authenticationMiddlewareImpl func() mux.MiddlewareFunc
	generateTokenImpl            func(username string) string
}

func (m *mockJwtService) AuthenticationMiddleware() mux.MiddlewareFunc {
	return m.authenticationMiddlewareImpl()
}

func (m *mockJwtService) GenerateToken(username string) string {
	return m.generateTokenImpl(username)
}

func Test_serviceImpl_Login_nilCredentials(t *testing.T) {
	repository := &mockRepository{findCredentialsByUsernameImpl: func(username string) (*u.Credentials, error) { return nil, nil }}
	service := serviceImpl{repository: repository}

	token, err := service.Login("unknown-user", "PaSsW0rD!")
	if token != "" {
		t.Errorf("Unexpected token: %v", token)
	}
	if err == nil {
		t.Error("Expected an error, got none")
	}
}

func Test_serviceImpl_Login_errorFindingCredentials(t *testing.T) {
	repository := &mockRepository{findCredentialsByUsernameImpl: func(username string) (*u.Credentials, error) { return nil, errors.New("Error finding credentials") }}
	service := serviceImpl{repository: repository}

	token, err := service.Login("some-user", "PaSsW0rD!")
	if token != "" {
		t.Errorf("Unexpected token: %v", token)
	}
	if err == nil {
		t.Error("Expected an error, got none")
	}
}

func Test_serviceImpl_Login_wrongPassword(t *testing.T) {
	user := &u.User{Id: 2, Username: "John Doe"}
	credentials := &u.Credentials{
		User:     user,
		Password: "$2a$10$bS.HuGI.l5pFEjfjDIjB2.3t9h62kRSi3exUTBhbs6vqrJouNTDh2",
		Roles:    []string{"ROLE_USER"},
	}

	repository := &mockRepository{findCredentialsByUsernameImpl: func(username string) (*u.Credentials, error) { return credentials, nil }}
	service := serviceImpl{repository: repository}

	token, err := service.Login(user.Username, "wrong-password")
	if token != "" {
		t.Errorf("Unexpected token: %v", token)
	}
	if err == nil {
		t.Error("Expected an error, got none")
	}
}
