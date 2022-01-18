package security

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"go-example/user"
	"go-example/util"
	"net/http"
	"strings"
)

type AuthenticationManager interface {
	AuthenticationMiddleware() mux.MiddlewareFunc
}

type authenticationManagerImpl struct {
	config         *util.Config
	userRepository user.Repository
}

func NewAuthenticationManager(config *util.Config, userRepository user.Repository) AuthenticationManager {
	return &authenticationManagerImpl{config: config, userRepository: userRepository}
}

func (m *authenticationManagerImpl) AuthenticationMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := m.extractUsernameFromRequest(r)
			if username != "" {
				//TODO load credentials and set them to the request context
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (m *authenticationManagerImpl) extractUsernameFromRequest(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		claims := &jwt.RegisteredClaims{}

		_, err := m.parseJwtToken(authorization[7:], claims)
		if err != nil {
			return ""
		}

		//TODO check expiration
		username := claims.Subject
		return username
	}

	return ""
}

func (m *authenticationManagerImpl) parseJwtToken(token string, claims jwt.Claims) (*jwt.Token, error) {
	validMethods := []string{jwt.SigningMethodHS512.Alg()}
	parser := jwt.NewParser(jwt.WithValidMethods(validMethods))
	return parser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.config.JwtSecret), nil
	})
}
