package user

import (
	"testing"
)

func Test_serviceImpl_Login_withSuccess(t *testing.T) {
	user := &User{Id: 2, Username: "John Doe"}
	credentials := &Credentials{
		User:     user,
		Password: "$2a$10$bS.HuGI.l5pFEjfjDIjB2.3t9h62kRSi3exUTBhbs6vqrJouNTDh2",
		Roles:    []string{"ROLE_USER"},
	}

	repository := &mockRepository{findCredentialsByUsernameImpl: func(username string) (*Credentials, error) { return credentials, nil }}

	service := serviceImpl{repository: repository}
	token, err := service.Login(user.Username, "PaSsW0rD!")
	if err != nil {
		t.Errorf("Expected no errors, got %v", err)
	}
	if token == "" {
		t.Error("Expected a valid token, got none")
	}
}

type mockRepository struct {
	findAllImpl                   func() ([]*User, error)
	findCredentialsByUsernameImpl func(username string) (*Credentials, error)
}

func (r *mockRepository) FindAll() ([]*User, error) {
	return r.findAllImpl()
}

func (r *mockRepository) FindCredentialsByUsername(username string) (*Credentials, error) {
	return r.findCredentialsByUsernameImpl(username)
}

func Test_serviceImpl_Login_nilCredentials(t *testing.T) {

}

func Test_serviceImpl_Login_errorFindingCredentials(t *testing.T) {

}

func Test_serviceImpl_Login_wrongPassword(t *testing.T) {

}
