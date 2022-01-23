package security

import (
	"go-example/user"
	"go-example/util"
	"net/http"
)

func AssertRole(request *http.Request, role string) (*user.Credentials, *util.HttpError) {
	credentials, ok := request.Context().Value("credentials").(*user.Credentials)
	if !ok {
		return nil, util.NewHttpError("Unauthorized access", http.StatusUnauthorized)
	}
	if !util.ContainsString(credentials.Roles, role) {
		return nil, util.NewHttpError("Unauthorized access", http.StatusForbidden)
	}

	return credentials, nil
}
