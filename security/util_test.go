package security

import (
	"context"
	"github.com/Hikarikun92/go-example/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_AssertRole_withoutCredentials(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	credentials, err := AssertRole(request, ROLE_USER)

	if credentials != nil {
		t.Errorf("Expected no credentials, got %v", credentials)
	}
	if err == nil {
		t.Error("Expected an error, got none")
	}
	if err.Status != http.StatusUnauthorized {
		t.Errorf("Expected status %v, got %v", http.StatusUnauthorized, err.Status)
	}
}

func Test_AssertRole_withoutRole(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	credentials := &user.Credentials{
		User:     &user.User{Id: 2, Username: "John Doe"},
		Password: "some-password",
		Roles:    []string{ROLE_USER},
	}
	ctx := context.WithValue(request.Context(), "credentials", credentials)
	request = request.WithContext(ctx)

	credentials, err := AssertRole(request, ROLE_ADMIN)

	if credentials != nil {
		t.Errorf("Expected no credentials, got %v", credentials)
	}
	if err == nil {
		t.Error("Expected an error, got none")
	}
	if err.Status != http.StatusForbidden {
		t.Errorf("Expected status %v, got %v", http.StatusForbidden, err.Status)
	}
}

func Test_AssertRole_correctRole(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	expectedCredentials := &user.Credentials{
		User:     &user.User{Id: 2, Username: "John Doe"},
		Password: "some-password",
		Roles:    []string{ROLE_USER},
	}
	ctx := context.WithValue(request.Context(), "credentials", expectedCredentials)
	request = request.WithContext(ctx)

	credentials, err := AssertRole(request, ROLE_USER)
	if credentials != expectedCredentials {
		t.Errorf("Expected %v, got %v", expectedCredentials, credentials)
	}
	if err != nil {
		t.Error("Unexpected error", err)
	}
}
