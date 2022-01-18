package security

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func AuthenticationMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			if strings.HasPrefix(authorization, "Bearer ") {
				//TODO validate and parse the token, extract the user credentials and set them to the request context
				//token := authorization[7:]
			}

			next.ServeHTTP(w, r)
		})
	}
}
