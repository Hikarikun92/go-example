package security

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"go-example/util"
	"net/http"
	"strings"
)

func AuthenticationMiddleware(config *util.Config) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := extractUsernameFromRequest(config, r)
			if username != "" {
				//TODO load credentials and set them to the request context
			}

			next.ServeHTTP(w, r)
		})
	}
}

func extractUsernameFromRequest(config *util.Config, r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		token, err := parseJwtToken(config, authorization[7:])
		if err != nil {
			return ""
		}

		//TODO check expiration
		username := token.Claims.(*jwt.RegisteredClaims).Subject
		return username
	}

	return ""
}

func parseJwtToken(config *util.Config, token string) (*jwt.Token, error) {
	validMethods := []string{jwt.SigningMethodHS512.Alg()}
	parser := jwt.NewParser(jwt.WithValidMethods(validMethods))
	return parser.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
}
