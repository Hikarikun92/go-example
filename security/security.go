package security

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"go-example/user"
	"go-example/util"
	"net/http"
	"strings"
	"time"
)

type AuthenticationManager interface {
	AuthenticationMiddleware() mux.MiddlewareFunc
	GenerateToken(username string) string
}

type authenticationManagerImpl struct {
	config         *util.Config
	userRepository user.Repository
}

func NewAuthenticationManager(config *util.Config, userRepository user.Repository) AuthenticationManager {
	return &authenticationManagerImpl{config: config, userRepository: userRepository}
}

//AuthenticationMiddleware returns a mux.MiddlewareFunc that will intercept each request, check if a JWT Bearer Authorization
//header is present and, if so, attempt to parse and validate it. If it is valid, the username corresponding to the "Subject"
//claim of the token is then loaded from the database and set to the current request context.
func (m *authenticationManagerImpl) AuthenticationMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := m.extractUsernameFromRequest(r)
			if username != "" {
				credentials, err := m.userRepository.FindCredentialsByUsername(username)
				if credentials != nil && err == nil {
					//Put the credentials in the request context; you can access it later with r.Context().Value("credentials")
					ctx := context.WithValue(r.Context(), "credentials", credentials)
					r = r.WithContext(ctx)
				}
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
		//In case of error, ignore it and return an empty username
		if err != nil {
			return ""
		}

		return claims.Subject
	}

	return ""
}

func (m *authenticationManagerImpl) parseJwtToken(token string, claims jwt.Claims) (*jwt.Token, error) {
	//Parse using HmacSHA512
	validMethods := []string{jwt.SigningMethodHS512.Alg()}
	parser := jwt.NewParser(jwt.WithValidMethods(validMethods))

	//Try to parse the token using the specified configurations, or return an error
	return parser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		//Return the signing key
		return m.config.JwtSecret, nil
	})
}

//GenerateToken generates a JWT token setting the given username as Subject, using HmacSHA512 signature and set to expire
//in a time period defined in the util.Config struct.
func (m *authenticationManagerImpl) GenerateToken(username string) string {
	now := time.Now()
	expiration := now.Add(time.Millisecond * m.config.JwtExpirationMs)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiration),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(m.config.JwtSecret)
	return token
}
