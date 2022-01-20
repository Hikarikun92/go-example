package rest

import (
	"testing"
)

func Test_controllerImpl_Login_withSuccess(t *testing.T) {
	panic("Implement me")
}

type mockFacade struct {
	loginImpl func(dto *LoginDto) (string, error)
}

func (f *mockFacade) Login(dto *LoginDto) (string, error) {
	return f.loginImpl(dto)
}

func Test_controllerImpl_Login_invalidJson(t *testing.T) {
	panic("Implement me")
}

func Test_controllerImpl_Login_loginError(t *testing.T) {
	panic("Implement me")
}
