package utils

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct {
	JwtMiddleware *jwtmiddleware.JWTMiddleware
	SecretKey     string "cliburn"
	SigningMethod jwt.SigningMethod
	Token         string
}

func (at *Auth) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (at *Auth) CreateToken(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	// Embed Admin information to `token` and Expire in 10 mins
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(at.SecretKey))
	if err != nil {
		return "Error occured"
	}
	return tokenString
}

func (at *Auth) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func (at *Auth) Init() {
	at.JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(at.SecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
